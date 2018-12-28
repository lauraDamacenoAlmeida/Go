package main
import(
	"fmt"
	"os"
	"net/http"
	"time"
)
const monitoramentos = 3
const delay = 5

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

//fazer requisições HTTP, para fazer o monitoramento do site escolhido e define o tempo para testar
func iniciarMonitoramento(){
	fmt.Println("Monitorando....")

	for i:=0; i<monitoramentos;i++{
	site := escolhaDoSite(i,getSites())
	fmt.Println("Testando o site ",site)
	testaSite(site)
	time.Sleep(delay * time.Second)
	}
}

func testaSite(site string){
	resp,_ := http.Get(site)
	if resp.StatusCode == 200{
		fmt.Println("Site foi carregado com sucesso")
	}else{
		fmt.Println("Chama o SAMUUU  \nStatus Code",resp.StatusCode)
	}
}

func escolhaDoSite(i int,sites[]string)string{
	return sites[i]
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
