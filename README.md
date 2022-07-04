
# Brazil Web Scrapy using Golang

``` Go | gocolly | Web Crawler | Web Scrap | Open Source 💙  ```

This is a small project to get noticies from [globo G1 portal](https://g1.globo.com/) using [gocolly](https://github.com/gocolly/colly)

## Aplication Steps
- Open search page
- Get links
- Clean links
- Open notice link page
- Extract informations

```go
type NoticeModule struct {
	Title   string
	Summary string
	Date    string
	Link    string
	Body    string
}
```
## Output
```
[
    {
        "Title": "VÍDEO: Trabalhadora rural se emociona ao receber notícia de aposentadoria e viraliza nas redes sociais",
        "Summary": "Ana Maria é maranhense e trabalhou por quase 50 anos na roça. Vídeo em que chora abraçando o advogado tem quase 130 mil visualizações.",
        "Date": "2022-07-01T14:34:16.060Z",
        "Link": "https://g1.globo.com/ma/maranhao/vem-ver-pequeno/noticia/2022/07/01/video-trabalhadora-rural-se-emociona-ao-receber-noticia-de-aposentadoria-e-viraliza-nas-redes-sociais.ghtml",
        "Body": "Essa semana, o vídeo de Ana Maria Mendes, de 55 anos, recebendo a notícia de sua aposentadoria após 49 anos de trabalho na roça, viralizou em todo o país. Nas imagens compartilhadas em uma rede social pelo seu advogado, Alexandre Pereira, Ana Maria não segurou a emoção ao saber que havia conquistado a tão sonhada aposentadoria. O caso aconteceu em Apicum-Açu, município localizado a 301 km de São Luís, onde Maria vive em uma comunidade quilombola.  O vídeo, gravado por câmeras de segurança do escritório de Alexandre, mostra Ana acompanhada da filha, Alicelea Oliveira, de 34 anos, que era responsável por receber as informações e respostas do processo administrativo. Após Alexandre entregar uma sacola com um presente de aniversário e dar a notícia da aposentadoria, Ana vai às lágrimas e abraça o advogado. A cena comovente acumula quase 130 mil visualizações em uma rede social (veja o vídeo acima).  Desde os seis anos de idade, Ana trabalhava com atividades ligadas ao campo, como plantações de alimentos, produção de farinha e capinagem. Analfabeta, ela começou a trabalhar muito cedo para ajudar financeiramente a mãe. Após criar nove filhos sozinha e trabalhar 49 anos na roça, Ana desejava um descanso que só poderia ser concedido pela sonhada aposentadoria.  Para Alicelea, a notícia da aposentadoria emocionou toda a família. Segundo ela, a mãe é uma guerreira que batalhou a vida inteira para cuidar dos filhos.  Em entrevista ao g1 Maranhão, o advogado Alexandre conta que, após a trabalhadora não conseguir obter o benefício por meio do Sindicato dos Trabalhadores Rurais, onde realizava contribuições mensais como lavradora, a filha de Ana o procurou e explicou a situação da mãe.  Segundo Alexandre, quando Maria “quebrou o protocolo” e o abraçou, ele ficou sem reação. “Fiquei muito surpreso porque no interior o advogado tem uma certa imagem que é inacessível, tem uma figura de autoridade, e geralmente as pessoas têm um grande respeito. Quando ela quebrou o protocolo, fiquei sem reação. Mas depois que recebi aquele abraço de gratidão, percebi que eu estava no caminho certo, porque ela poderia ser só mais uma cliente, mas, para ela, sou uma pessoa que contribuiu para mudar uma parcela da vida dela”, disse.  De início, o advogado compartilhou o vídeo apenas com seus dois sócios da cidade de São Luís, mas após ser incentivado por eles e outros amigos, resolveu publicar o registro em sua rede social. Para Alexandre, a viralização do vídeo contribui como algo positivo e motivacional no cenário da advocacia.  *Matéria realizada sob supervisão de Rafael Cardoso, g1 MA."
    },

]
```
