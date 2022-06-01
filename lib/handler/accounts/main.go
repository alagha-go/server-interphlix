package accounts

import "interphlix/lib/accounts"


func Main() {
	accounts.Main()
}

func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}