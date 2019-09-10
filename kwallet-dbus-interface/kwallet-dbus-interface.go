package main

// NOTE: Some of the calls take int, some int64... for no apparent reason it seems

import (
	"C"
	"errors"
	"fmt"
	"github.com/godbus/dbus"
	"log"
)

const (
	applicationName string = "pidgin-kwallet-plugin"
	folderName      string = "pidgin"
	dbusName        string = "pidgin.pidgin-kwallet-plugin"
)

func getHandle(wallet string) (int64, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return -1, err
	}
	handle := int64(-1)
	err = conn.Object("org.kde.kwalletd5", "/modules/kwalletd5").Call("org.kde.KWallet.open", 0, wallet, int64(0), applicationName).Store(&handle)
	if err != nil {
		return int64(-1), err
	}
	if handle == int64(-1) {
		return int64(-1), errors.New("Can not aquire handle. Do I have the correct permission?")
	}
	return handle, nil
}

func closeHandle(handle int64) error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	var returnCode int
	err = conn.Object("org.kde.kwalletd5", "/modules/kwalletd5").Call("org.kde.KWallet.close", 0, handle, applicationName).Store(&returnCode)
	if err != nil {
		return err
	}
	if returnCode == -1 {
		return errors.New("Can not close handle.")
	}
	return nil
}

//export CheckPidginDir
func CheckPidginDir(wallet *C.char) {
	handle, err := getHandle(C.GoString(wallet))
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Println(err)
		return
	}
	defer closeHandle(handle)
	var folderExists bool
	err = conn.Object("org.kde.kwalletd5", "/modules/kwalletd5").Call("org.kde.KWallet.hasFolder", 0, int(handle), folderName, applicationName).Store(&folderExists)
	if err != nil {
		log.Println(err)
		return
	}
	if !folderExists {
		var creationSuccessful bool
		err = conn.Object("org.kde.kwalletd5", "/modules/kwalletd5").Call("org.kde.KWallet.createFolder", 0, int(handle), folderName, applicationName).Store(&creationSuccessful)
		if err != nil {
			log.Println(err)
			return
		}
		if !creationSuccessful {
			fmt.Println("Could not create folder!")
		}
	}
}

//export GetPassword
func GetPassword(wallet, account, protocol *C.char) *C.char {
	fmt.Println("Getting password for account", C.GoString(account), C.GoString(protocol))
	handle, err := getHandle(C.GoString(wallet))
	if err != nil {
		log.Println(err)
		return nil
	}
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer closeHandle(handle)
	var password string
	err = conn.Object("org.kde.kwalletd5", "/modules/kwalletd5").Call("org.kde.KWallet.readPassword", 0, int(handle), folderName, fmt.Sprint(C.GoString(account), ";", C.GoString(protocol)), applicationName).Store(&password)
	if err != nil {
		log.Println(err)
		return nil
	}
	return C.CString(password)
}

//export SetPassword
func SetPassword(wallet, account, protocol, password *C.char) bool {
	fmt.Println("Saving new password for account", C.GoString(account))
	handle, err := getHandle(C.GoString(wallet))
	if err != nil {
		log.Println(err)
		return false
	}
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Println(err)
		return false
	}
	defer closeHandle(handle)
	var returnCode int
	err = conn.Object("org.kde.kwalletd5", "/modules/kwalletd5").Call("org.kde.KWallet.writePassword", 0, int(handle), folderName, fmt.Sprint(C.GoString(account), ";", C.GoString(protocol)), C.GoString(password), applicationName).Store(&returnCode)
	if err != nil {
		log.Println(err)
		return false
	}
	return returnCode != -1
}

func main() {}
