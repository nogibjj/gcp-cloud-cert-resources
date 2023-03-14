#!/usr/bin/env bash
source /home/codespace/venv/bin/activate
#append it to bash so every shell launches with it 
echo 'source /home/codespace/venv/bin/activate' >> ~/.bashrc
make install-tensorflow

## install rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y