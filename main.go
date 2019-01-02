package main
import(
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
	"io"
	"strings"
	"io/ioutil"
)
const delay = 5

func main(){
	registraLog("www.google.com",false)
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
		imprimeLogs()
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
	site := lerSites()
	for i:=0; i<len(site);i++{
	fmt.Println("Testando o site ",site[i])
	testaSite(site[i])
	time.Sleep(delay * time.Second)
	}
}

func testaSite(site string){
	resp,err := http.Get(site)
	if err != nil{
		fmt.Println("Ops!! Ocorreu um erro: ",err)
	}
	if resp.StatusCode == 200{
		fmt.Println("Site foi carregado com sucesso")
		registraLog(site,true)
	}else{
		fmt.Println("Chama o SAMUUU  \nStatus Code",resp.StatusCode)
		registraLog(site,false)
	}
}


func getSites()[]string{
	//sites := []string{"https://random-status-code.herokuapp.com/","https://www.meetup.com/pt-BR/bluetalks/events","https://www.google.com/"}
	sites := lerSites()
	return sites
}
func lerSites()[]string{
	var sites [] string
	arquivo, err := os.Open("sites.txt")
	if err != nil{
		fmt.Println("Ops!! Ocorreu um erro: ",err)
	}
	leitor := bufio.NewReader(arquivo)
	for{
	linha,err:= leitor.ReadString('\n')
	linha = strings.TrimSpace(linha)
	if err == io.EOF{
		break;
	}
	sites = append(sites,linha)
	}

	arquivo.Close()
	return sites
}
func registraLog(site string, status bool){
	arq,err := os.OpenFile("logs.txt",os.O_RDWR| os.O_CREATE| os.O_APPEND,0666)
	if err != nil{
		fmt.Println(err)
	}
	online := "No ar"
	if(status != true){
		online = "Site está fora do ar"
	}
	//para representar os dias e as horas atuais, no GO eh feito através dos numeros especificos
	arq.WriteString(time.Now().Format("02/01/2006 15:04:05")+" - Site: "+ site + " - status: " + online +"\n")
	arq.Close()
}

func imprimeLogs(){
	arq,err := ioutil.ReadFile("logs.txt")

	if err != nil{
		fmt.Println("OPSS, arquivo nao encontrado, favor verificar")
	}
	fmt.Println(string(arq))

}
