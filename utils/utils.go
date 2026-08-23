package utils
import (
	"error"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string) (string,error){
	return func()(string,error){
		if PasswordHash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost); err!=nil{
			return "",err
		}else{
			return password,nil
		}
	}()
}

func checkPasswordHash(password,hashstring string) bool{
	return func()bool{
		if err:=bcrypt.CompareHashAndPassword([](hash),byte[](password)); err!=nill{
			return err==nil
		}
	}()
}

