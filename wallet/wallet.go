package wallet

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "log"

    "golang.org/x/crypto/ripemd160"
)

const (
    checksumLength = 4
    //hexadecimal representation of 0
    version = byte(0x00)
)

type Wallet struct {
    //ecdsa = eliptical curve digital signiture algorithm
    PrivateKey ecdsa.PrivateKey
    PublicKey  []byte
}
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
    curve := elliptic.P256()

    private, err := ecdsa.GenerateKey(curve, rand.Reader)
    if err != nil {
        log.Panic(err)
    }

    pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

    return *private, pub
}
func PublicKeyHash(publicKey []byte) []byte {
    hashedPublicKey := sha256.Sum256(publicKey)

    hasher := ripemd160.New()
    _, err := hasher.Write(hashedPublicKey[:])
    if err != nil {
        log.Panic(err)
    }
    publicRipeMd := hasher.Sum(nil)

    return publicRipeMd
}

func Checksum(ripeMdHash []byte) []byte {
    firstHash := sha256.Sum256(ripeMdHash)
    secondHash := sha256.Sum256(firstHash[:])

    return secondHash[:checksumLength]
}


func (w *Wallet) Address() []byte {
    // Step 1/2
    pubHash := PublicKeyHash(w.PublicKey)
    //Step 3
    versionedHash := append([]byte{version}, pubHash...)
    //Step 4
    checksum := Checksum(versionedHash)
    //Step 5
    finalHash := append(versionedHash, checksum...)
    //Step 6
    address := base58Encode(finalHash)
    return address
}

