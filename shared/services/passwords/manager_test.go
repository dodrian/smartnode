package passwords

import (
    "io"
    "io/ioutil"
    "testing"
)


// Test password manager functionality
func TestPasswordManager(t *testing.T) {

    // Create temporary password path
    passwordPath, err := ioutil.TempDir("", "")
    if err != nil { t.Fatal(err) }
    passwordPath += "/password"

    // Create temporary input file
    input, err := ioutil.TempFile("", "")
    if err != nil { t.Fatal(err) }
    defer input.Close()

    // Write input to file
    io.WriteString(input, "foobarbaz" + "\n")
    input.Seek(0, io.SeekStart)

    // Initialise password manager
    passwordManager := NewPasswordManager(input, passwordPath)

    // Check if password exists
    passwordExists := passwordManager.PasswordExists()
    if passwordExists { t.Errorf("Incorrect password exists status: expected %t, got %t", false, passwordExists) }

    // Create password
    password, err := passwordManager.CreatePassword()
    if err != nil { t.Error(err) }
    if password != "foobarbaz" { t.Errorf("Incorrect created password: expected %s, got %s", "foobarbaz", password) }

    // Check if password exists
    passwordExists = passwordManager.PasswordExists()
    if !passwordExists { t.Errorf("Incorrect password exists status: expected %t, got %t", true, passwordExists) }

    // Get passphrase
    expectedPassphrase := "69a0dafe010dfa7ba062ea986bf94d20f16cf49e376e761bf679b6cc5b8cee6d"
    passphrase, err := passwordManager.GetPassphrase()
    if err != nil { t.Error(err) }
    if passphrase != expectedPassphrase { t.Errorf("Incorrect passphrase: expected %s, got %s", expectedPassphrase, passphrase) }

}
