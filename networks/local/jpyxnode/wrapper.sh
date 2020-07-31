#!/usr/bin/env sh

BINARY=/eurxd/linux/${BINARY:-eurxd}
echo "binary: ${BINARY}"
ID=${ID:-0}
LOG=${LOG:-eurxd.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'eurxd' E.g.: -e BINARY=eurxd_my_test_version"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export EURXDHOME="/eurxd/node${ID}/eurxd"

if [ -d "$(dirname "${EURXDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${EURXDHOME}" "$@" | tee "${EURXDHOME}/${LOG}"
else
  "${BINARY}" --home "${EURXDHOME}" "$@"
fi
