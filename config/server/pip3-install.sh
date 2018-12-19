#!/bin/sh

source /config/ENV.sh

if [[ $# == 0 ]]; then
	cat <<-EOS
	usage: $0 [packages...] 

	This script invokes two commands
	  - 'pip3 install [packages...]'
	  - 'pip3 freeze > ${PY_REQUIREMENTS_PATH}'
	EOS
else
	pip3 install $@
	pip3 freeze > $PY_REQUIREMENTS_PATH
fi
