# GnuPG

## Master Key

```bash
$ gpg --full-gen-key --expert
gpg (GnuPG) 2.2.35; Copyright (C) 2022 g10 Code GmbH
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Please select what kind of key you want:
   (1) RSA and RSA (default)
   (2) DSA and Elgamal
   (3) DSA (sign only)
   (4) RSA (sign only)
   (7) DSA (set your own capabilities)
   (8) RSA (set your own capabilities)
   (9) ECC and ECC
  (10) ECC (sign only)
  (11) ECC (set your own capabilities)
  (13) Existing key
  (14) Existing key from card
Your selection? 11

Possible actions for a ECDSA/EdDSA key: Sign Certify Authenticate
Current allowed actions: Sign Certify

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? S

Possible actions for a ECDSA/EdDSA key: Sign Certify Authenticate
Current allowed actions: Certify

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? q
Please select which elliptic curve you want:
   (1) Curve 25519
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
Your selection? 1
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 3y
Key expires at Mon 07 Jul 2025 12:09:00 AM IST
Is this correct? (y/N) y

GnuPG needs to construct a user ID to identify your key.

Real name: Murtaza Udaipurwala
Email address: murtaza@murtazau.xyz
Comment:
You selected this USER-ID:
    "Murtaza Udaipurwala <murtaza@murtazau.xyz>"

Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? O
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.
gpg: directory '/tmp/gpg/openpgp-revocs.d' created
gpg: revocation certificate stored as '/tmp/gpg/openpgp-revocs.d/7171ACBD6BC95FA386779BBB9240881A0950D364.rev'
public and secret key created and signed.

pub   ed25519 2022-07-07 [C] [expires: 2025-07-06]
      7171ACBD6BC95FA386779BBB9240881A0950D364
uid                      Murtaza Udaipurwala <murtaza@murtazau.xyz>
```

## Get `keyid`

```bash
$ gpg -k --keyid-format LONG
/tmp/gpg/pubring.kbx
--------------------
pub   ed25519/9240881A0950D364 2022-07-07 [C] [expires: 2025-07-06]
      7171ACBD6BC95FA386779BBB9240881A0950D364
uid                 [ultimate] Murtaza Udaipurwala <murtaza@murtazau.xyz>
```

Here, keyid = 9240881A0950D364

## Sub-keys

1. Sub-key with signing capability

```bash
$ gpg --edit-key --expert 9240881A0950D364
gpg (GnuPG) 2.2.35; Copyright (C) 2022 g10 Code GmbH
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Secret key is available.

sec  ed25519/9240881A0950D364
     created: 2022-07-07  expires: 2025-07-06  usage: C
     trust: ultimate      validity: ultimate
[ultimate] (1). Murtaza Udaipurwala <murtaza@murtazau.xyz>

gpg> addkey
Please select what kind of key you want:
   (3) DSA (sign only)
   (4) RSA (sign only)
   (5) Elgamal (encrypt only)
   (6) RSA (encrypt only)
   (7) DSA (set your own capabilities)
   (8) RSA (set your own capabilities)
  (10) ECC (sign only)
  (11) ECC (set your own capabilities)
  (12) ECC (encrypt only)
  (13) Existing key
  (14) Existing key from card
Your selection? 11

Possible actions for a ECDSA/EdDSA key: Sign Authenticate
Current allowed actions: Sign

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? q
Please select which elliptic curve you want:
   (1) Curve 25519
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
Your selection? 1
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 1y
Key expires at Sat 08 Jul 2023 12:10:05 AM IST
Is this correct? (y/N) y
Really create? (y/N) y
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.

sec  ed25519/9240881A0950D364
     created: 2022-07-07  expires: 2025-07-06  usage: C
     trust: ultimate      validity: ultimate
ssb  ed25519/5488588C540A42EC
     created: 2022-07-07  expires: 2023-07-07  usage: S
[ultimate] (1). Murtaza Udaipurwala <murtaza@murtazau.xyz>
```

2. Sub-key with authentication capability

```bash
gpg> addkey
Please select what kind of key you want:
   (3) DSA (sign only)
   (4) RSA (sign only)
   (5) Elgamal (encrypt only)
   (6) RSA (encrypt only)
   (7) DSA (set your own capabilities)
   (8) RSA (set your own capabilities)
  (10) ECC (sign only)
  (11) ECC (set your own capabilities)
  (12) ECC (encrypt only)
  (13) Existing key
  (14) Existing key from card
Your selection? 11

Possible actions for a ECDSA/EdDSA key: Sign Authenticate
Current allowed actions: Sign

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? S

Possible actions for a ECDSA/EdDSA key: Sign Authenticate
Current allowed actions:

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? A

Possible actions for a ECDSA/EdDSA key: Sign Authenticate
Current allowed actions: Authenticate

   (S) Toggle the sign capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? q
Please select which elliptic curve you want:
   (1) Curve 25519
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
Your selection? 1
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 1y
Key expires at Sat 08 Jul 2023 12:10:23 AM IST
Is this correct? (y/N) y
Really create? (y/N) y
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.

sec  ed25519/9240881A0950D364
     created: 2022-07-07  expires: 2025-07-06  usage: C
     trust: ultimate      validity: ultimate
ssb  ed25519/5488588C540A42EC
     created: 2022-07-07  expires: 2023-07-07  usage: S
ssb  ed25519/03AFC7671C6D0358
     created: 2022-07-07  expires: 2023-07-07  usage: A
[ultimate] (1). Murtaza Udaipurwala <murtaza@murtazau.xyz>
```

3. Lastly, sub-key with encryption capability

```bash
gpg> addkey
Please select what kind of key you want:
   (3) DSA (sign only)
   (4) RSA (sign only)
   (5) Elgamal (encrypt only)
   (6) RSA (encrypt only)
   (7) DSA (set your own capabilities)
   (8) RSA (set your own capabilities)
  (10) ECC (sign only)
  (11) ECC (set your own capabilities)
  (12) ECC (encrypt only)
  (13) Existing key
  (14) Existing key from card
Your selection? 8

Possible actions for a RSA key: Sign Encrypt Authenticate
Current allowed actions: Sign Encrypt

   (S) Toggle the sign capability
   (E) Toggle the encrypt capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? S

Possible actions for a RSA key: Sign Encrypt Authenticate
Current allowed actions: Encrypt

   (S) Toggle the sign capability
   (E) Toggle the encrypt capability
   (A) Toggle the authenticate capability
   (Q) Finished

Your selection? q
RSA keys may be between 1024 and 4096 bits long.
What keysize do you want? (3072) 4096
Requested keysize is 4096 bits
Please specify how long the key should be valid.
         0 = key does not expire
      <n>  = key expires in n days
      <n>w = key expires in n weeks
      <n>m = key expires in n months
      <n>y = key expires in n years
Key is valid for? (0) 1y
Key expires at Sat 08 Jul 2023 12:10:43 AM IST
Is this correct? (y/N) y
Really create? (y/N) y
We need to generate a lot of random bytes. It is a good idea to perform
some other action (type on the keyboard, move the mouse, utilize the
disks) during the prime generation; this gives the random number
generator a better chance to gain enough entropy.

sec  ed25519/9240881A0950D364
     created: 2022-07-07  expires: 2025-07-06  usage: C
     trust: ultimate      validity: ultimate
ssb  ed25519/5488588C540A42EC
     created: 2022-07-07  expires: 2023-07-07  usage: S
ssb  ed25519/03AFC7671C6D0358
     created: 2022-07-07  expires: 2023-07-07  usage: A
ssb  rsa4096/F8C47936EF059774
     created: 2022-07-07  expires: 2023-07-07  usage: E
[ultimate] (1). Murtaza Udaipurwala <murtaza@murtazau.xyz>

gpg> save
```

## Key note

* The master key and revocation certificate should be encrypted with
  symmetric encryption and stored safely in an encrypted volume on
  separate USB thumb drives.
* These thumb drives should eventually be stored in a secret safe which
  only you have access to. Hopefully, our master keys shall never see
  the light again.
* The sub-keys are meant for daily use.
* In case your sub-keys get compromised, rush to the safe, get the USB
  thumb drive out, and revoke the sub keys. This preserves our identify
  in the web of trust.
* You’ll revoke those sub-keys using the master key, not the revocation
  certificate.
* The revocation certificate is used to revoke the master key only
  (that's the reason we are storing them on a separate thumb drive).

## References

* <https://gutier.io/post/security-generate-strongest-secure-gpg-key/>
* <https://alexcabal.com/creating-the-perfect-gpg-keypair>
