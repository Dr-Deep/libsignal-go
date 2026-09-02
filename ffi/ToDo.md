

/* Key-Generierung (Install-/Link-Zeit)
IdentityKeyPair.generate()
generateRegistrationId()
PrivateKey.generate() / PublicKey
generatePreKey(keyId)
generateSignedPreKey(identityKeyPair, keyId)
generateKyberPreKey(...)          // PQ
*/

/*
   ProvisioningConnection: für linked device
   ChatConnection: persistent
   ChatListener: gibt uns neue messages
   8. Bei neuer Session → SessionBuilder + PreKeyBundle
   9. Entschlüsseln mit SessionCipher.decrypt / decryptPreKeySignalMessage
*/
