#!/usr/bin/env bash
# The stable signature of the live-evals tool (D-1208).
#
# macOS records the access to a removable volume against the designated
# requirement of a program. The Go linker signs each build ad hoc, and
# that requirement is the hash of the build. So each new build asked
# again for the access (D-1178). A signature with one certificate gives
# each build the same requirement, and the grant stays.
#
#   scripts/live-evals-sign.sh setup        make the key, one time (make live-evals-sign-setup)
#   scripts/live-evals-sign.sh sign FILE    sign one build (the tick calls it)
#   scripts/live-evals-sign.sh requirement  print the requirement of the key
#
# The key lives in a keychain of its own in LIVE_EVALS_SECRETS, with its
# password in a file of mode 600 beside it. The sandbox of a session
# denies reads under HOME, so a session can not read either file. A
# session can still use an unlocked keychain through securityd. So the
# sign step unlocks the keychain, signs, and locks it again before any
# session starts.
set -uo pipefail

SECRETS=${LIVE_EVALS_SECRETS:-$HOME/.config/decktome-live-evals}
KEYCHAIN=$SECRETS/sign.keychain-db
PASS_FILE=$SECRETS/sign-keychain-pass
ID_FILE=$SECRETS/sign-identity
IDENT=com.decktome.live-evals
CN=decktome-live-evals
# The openssl of macOS. A conda or Homebrew openssl 3 writes a PKCS#12
# file that security(1) can not read.
OPENSSL=/usr/bin/openssl

die() { echo "live-evals-sign: $*" >&2; exit 1; }

requirement() { # sha1
  printf 'identifier "%s" and certificate leaf = H"%s"' "$IDENT" "$1"
}

# codesign finds an identity only in a keychain of the search list. The
# keychain joins the list for the sign alone, because a locked keychain
# on the list can make other apps ask for its password.
read_list() { # sets LIST to the search list of the user
  local line
  LIST=()
  while IFS= read -r line; do
    line=${line#"${line%%[![:space:]]*}"}
    line=${line#\"}
    line=${line%\"}
    [ -n "$line" ] && LIST+=("$line")
  done < <(security list-keychains -d user)
}

with_keychain() { # command...
  local rc
  read_list
  security list-keychains -d user -s "${LIST[@]}" "$KEYCHAIN" || die "can not add $KEYCHAIN to the search list"
  "$@"
  rc=$?
  security list-keychains -d user -s "${LIST[@]}" || die "can not restore the search list of keychains"
  return "$rc"
}

setup() {
  [ -e "$KEYCHAIN" ] && die "$KEYCHAIN exists. Delete it and $PASS_FILE first to make a new key. The new key asks for the grant one more time."
  mkdir -p "$SECRETS" || die "can not make $SECRETS"
  chmod 700 "$SECRETS" || die "can not set the mode of $SECRETS"
  local pass sha
  TMP=$(mktemp -d) || die "can not make a temporary folder"
  trap 'rm -rf "$TMP"' EXIT
  local tmp=$TMP
  cat > "$tmp/req.cnf" <<EOF
[req]
distinguished_name=dn
prompt=no
x509_extensions=v3
[dn]
CN=$CN
[v3]
basicConstraints=critical,CA:false
keyUsage=critical,digitalSignature
extendedKeyUsage=critical,codeSigning
EOF
  "$OPENSSL" req -x509 -newkey rsa:2048 -nodes -days 3650 -config "$tmp/req.cnf" \
    -keyout "$tmp/key.pem" -out "$tmp/cert.pem" 2>/dev/null || die "openssl can not make the certificate"
  pass=$("$OPENSSL" rand -hex 32) || die "openssl can not make the password"
  "$OPENSSL" pkcs12 -export -inkey "$tmp/key.pem" -in "$tmp/cert.pem" -out "$tmp/id.p12" \
    -passout "pass:$pass" || die "openssl can not pack the key"
  sha=$("$OPENSSL" x509 -in "$tmp/cert.pem" -noout -fingerprint -sha1 | cut -d= -f2 | tr -d : | tr 'A-F' 'a-f')
  (umask 077 && printf '%s' "$pass" > "$PASS_FILE") || die "can not write $PASS_FILE"
  # A creation adds the keychain to the search list. The sign adds it
  # for its own call alone, so the setup keeps the list as it was.
  read_list
  security create-keychain -p "$pass" "$KEYCHAIN" || die "can not make $KEYCHAIN"
  security list-keychains -d user -s "${LIST[@]}" || die "can not take $KEYCHAIN off the search list"
  security set-keychain-settings -l -u -t 300 "$KEYCHAIN" || die "can not set the lock of $KEYCHAIN"
  security unlock-keychain -p "$pass" "$KEYCHAIN" || die "can not unlock $KEYCHAIN"
  security import "$tmp/id.p12" -k "$KEYCHAIN" -P "$pass" -T /usr/bin/codesign >/dev/null \
    || die "can not import the key"
  security set-key-partition-list -S apple-tool:,apple:,codesign: -s -k "$pass" "$KEYCHAIN" >/dev/null \
    || die "can not let codesign use the key"
  security lock-keychain "$KEYCHAIN" || die "can not lock $KEYCHAIN again. Lock it by hand."
  printf '%s\n' "$sha" > "$ID_FILE"
  chmod 600 "$KEYCHAIN" "$ID_FILE"
  echo "live-evals-sign: the key is ready. Each build now has this requirement:"
  echo "  $(requirement "$sha")"
  echo "The next tick asks for the grant of the volume one more time. Allow it on the Mac."
}

sign() { # file
  local file=$1 sha pass got want
  [ -f "$file" ] || die "no file $file"
  [ -s "$ID_FILE" ] || die "no $ID_FILE. Run make live-evals-sign-setup."
  sha=$(cat "$ID_FILE")
  local mode
  mode=$(stat -f %Lp "$PASS_FILE") || die "can not read the mode of $PASS_FILE"
  [ $(( 8#$mode & 8#077 )) = 0 ] || die "$PASS_FILE has mode $mode, and another account can read it. Run chmod 600 on it (D-1164)."
  pass=$(cat "$PASS_FILE") || die "can not read $PASS_FILE"
  security unlock-keychain -p "$pass" "$KEYCHAIN" || die "can not unlock $KEYCHAIN"
  # codesign writes a note on each success, so the tick logs its text
  # only on a failure.
  local out rc
  out=$(with_keychain codesign -f -s "$sha" -i "$IDENT" "$file" 2>&1)
  rc=$?
  security lock-keychain "$KEYCHAIN" || die "can not lock $KEYCHAIN again. Lock it by hand before the next tick."
  [ "$rc" = 0 ] || die "codesign did not sign $file: $out"
  got=$(codesign -d -r- "$file" 2>&1 | sed -n 's/^designated => //p')
  want=$(requirement "$sha")
  [ "$got" = "$want" ] || die "the requirement of $file is [$got], not [$want]"
}

case "${1:-}" in
  setup) setup ;;
  sign) [ $# = 2 ] || die "usage: sign FILE"; sign "$2" ;;
  requirement) [ -s "$ID_FILE" ] || die "no $ID_FILE"; requirement "$(cat "$ID_FILE")"; echo ;;
  *) die "usage: setup | sign FILE | requirement" ;;
esac
