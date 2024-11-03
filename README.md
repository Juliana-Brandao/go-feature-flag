# Feature Flag
Essa aplicação é um exemplo prático do uso do conceito de Feature Flag para aplicações em larga escala.

### Solução
Aplicações distribuídas e de larga escala podem ter centenas ou até milhares de instâncias em execução simultânea. O controle da ativação e desativação de feature flags precisa ser realizado de forma rápida e segura para permitir a ativação de novas funcionalidades para testes, mas também possibilitar a desativação imediata em caso de problemas.

A solução proposta é simples e eficiente. Ela consiste em um banco de dados central que mantém o status de cada feature flag, o qual é consultado a cada minuto por cada instância da aplicação. Além disso, a aplicação envia métricas para o Prometheus, permitindo que os desenvolvedores acompanhem o status de cada feature flag em todas as instâncias. Isso viabiliza um monitoramento preciso e em tempo real das funcionalidades em produção.
### Arquitetura
![Architecture](documentation/images/architecture.png)

### Executando
A aplicação está configurada para rodar com Docker e Docker Compose. Portanto, para vê-la funcionando, basta apenas entrar no diretório e executar o comando

`
docker-compose up --build
`
Após inicializar o Docker Compose, acesse o [Grafana clicando aqui](http://localhost:3000) use o `admin` como usuário e senha. Após o login importe o dashboard a partir do arquivo `dashboard.json` nesse repositório.



### Arquivos