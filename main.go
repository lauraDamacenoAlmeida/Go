package main

import "fmt"
import "os"
import "net/http"

func main(){
	for{
	exibeMenu()
	lerEntrada()
	}
}
func validaEntrada(entrada int){
	switch entrada {
	case 1:
		iniciarMonitoramento()
		break
	case 2:
		fmt.Println("Opcao 2")
	case 9:
		fmt.Println("Adeus...")
		os.Exit(0)
	default:
		fmt.Println("Opsss!!! Nao entendi este comando")
	}


}

func lerEntrada(){
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println(comando)
	validaEntrada(comando)
}

func exibeMenu(){
	fmt.Println("Digite uma opcao:")
	fmt.Println("1- Monitoramento \n2- Exibir logs \n9- Sair")
}

//fazer requisições HTTP, para fazer o monitoramento do site escolhido
func iniciarMonitoramento(){
	fmt.Println("Monitorando....")
	site := escolhaDoSite(getSites())
	testaSite(site)

}

func testaSite(site string){
	resp,_ := http.Get(site)
	if resp.StatusCode == 200{
		fmt.Println("Site",site,"foi carregado com sucesso")
	}else{
		fmt.Println("Site ",site)
		fmt.Println("Chama o SAMUUU  \nStatus Code",resp.StatusCode)
	}
}

func escolhaDoSite(sites[]string)string{

	for i,site:=range sites{
		fmt.Println(i,site)
		return site
	}
	return "https://random-status-code.herokuapp.com/"
}

func getSites()[]string{
	sites := []string{"https://random-status-code.herokuapp.com/","https://www.meetup.com/pt-BR/bluetalks/events","https://www.google.com/"}
	return sites
}

func setSites(){
	var site string
	fmt.Println("Digite o novo site")
	fmt.Scanf("%s", &site)
	fmt.Println(site)
	//sites = append(sites,site)
}
