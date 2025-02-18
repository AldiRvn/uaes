# Note

## JS Example

[playground](https://stackblitz.com/edit/node-lpvhzq?embed=1&file=index.js)

```node
var CryptoJS = require('crypto-js');

var decrypted = CryptoJS.AES.decrypt(
  'U2FsdGVkX1+Eapvx6VOPBY5Bx4sLKx+R3830AZOTu0Q=',
  '%4NmStvr4@NrVheI'
);
console.log(decrypted.toString(CryptoJS.enc.Utf8)); // Zokijda
```

## PY Example

Library: pycryptodomex

pycryptojs.py

```py
# source: https://stackoverflow.com/a/36780727
# pip install pycryptodomex
from Cryptodome import Random
from Cryptodome.Cipher import AES
import base64
from hashlib import md5

BLOCK_SIZE = 16


def pad(data):
    length = BLOCK_SIZE - (len(data) % BLOCK_SIZE)
    return data + (chr(length) * length).encode()


def unpad(data):
    return data[:-(data[-1] if type(data[-1]) == int else ord(data[-1]))]


def bytes_to_key(data, salt, output=48):
    # extended from https://gist.github.com/gsakkis/4546068
    assert len(salt) == 8, len(salt)
    data += salt
    key = md5(data).digest()
    final_key = key
    while len(final_key) < output:
        key = md5(key + data).digest()
        final_key += key
    return final_key[:output]


def encrypt(message, secret):
    salt = Random.new().read(8)
    key_iv = bytes_to_key(secret, salt, 32 + 16)
    key = key_iv[:32]
    iv = key_iv[32:]
    aes = AES.new(key, AES.MODE_CBC, iv)
    return base64.b64encode(b"Salted__" + salt + aes.encrypt(pad(message)))


def decrypt(encrypted, secret):
    encrypted = base64.b64decode(encrypted)
    assert encrypted[0:8] == b"Salted__"
    salt = encrypted[8:16]
    key_iv = bytes_to_key(secret, salt, 32 + 16)
    key = key_iv[:32]
    iv = key_iv[32:]
    aes = AES.new(key, AES.MODE_CBC, iv)
    return unpad(aes.decrypt(encrypted[16:]))
```

main.py

```py
import pycryptojs

secret = 'yourSecretEncryption'.encode()
encrypted = pycryptojs.encrypt(secret=secret, message='Pagi.'.encode())
print('Test encrypt result:')
print(encrypted.decode('utf'))

result = pycryptojs.decrypt(secret=secret, encrypted=encrypted)
print('Test decode result:')
print(result.decode('utf'))
```
