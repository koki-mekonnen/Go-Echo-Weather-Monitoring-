package configs

type ServerConfig struct{
	Address string
}


var Server ServerConfig

func Load(){
	Server.Address=":8000"
}
