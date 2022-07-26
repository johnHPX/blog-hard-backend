# Estrutura das Requisições e Respostas

request e response serão enviadas por JSON.

#### - _BODY_

exemplos de JSON a serem enviadoas pelo body's Request

```
## object

{
  "attribute: "value",
  "attribute": [{
    "attribute: "value",
  }]
}

## array

[
  {
    "attribute: "value",
  },
  {
    "attribute: "value",
  }
]

```

#### - _QUERIES_

passando dados pela URL.

/example?attribute=value&attribute=value

# CONCEITOS

  | Nome             | Descrição                                                                                        |
  | ---------------- | ------------------------------------------------------------------------------------------------ |
  | request          | diz se os dados em sua maioria serão enviados pelos 3 tipos de paramentros(body, queries ou url) |
  | type             | diz o tipo de estrutura que será enviada. exemplo: objeto, array...                              |
  | method           | diz o verbo http que a rota utiliza. exemplo (post, get, put e delete)                           |
  | toke is requered | diz se é necessario um token de acesso para utilização desse endpoint                            |
  | attribute name   | diz o nome do atributo                                                                           |
  | type value       | diz a tipagem do atributo                                                                        |
  | size             | diz o valoz maximo de caracteres que o atributo pode ter                                         |
  | is it requerid?  | diz se o atributo é obrigatorio                                                                  |
  | type send        | diz a maneira que o atributo será enviado. exemplo: body, queries ou url paraments               |
  | description      | a descrição do atributo                                                                          |
  | status           | o status do response. exemplo: 200, 201...                                                       |

  ```
    OBS: MID é um codigo curto que o frontEnd pode enviar a API, com o intuito de verificar se a requisição ocorreu, sendo bem sucessido ou mal sucessido.

    Dependendo do METHOD, ele pode ser enviado por BODY'S request ou QUERY'S request.
    EXEMPLO:  
    // body's request
     "mid": "ok"

    // query's request
    ? mid=ok

  ```

<hr>
<h1> USER Routes </h1>


## 1. /user/store

criando uma novo usuario.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | not               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `name`         | `string`   | `255` | `true`          | body paraments | nome do usuario                                  |
| `telephone`    | `string`   | `13`  | `true`          | body paraments | telefone do usuario                              |
| `nick`         | `string`   | `255` | `true`          | body paraments | nick do usuario                                  |
| `email`        | `string`   | `255` | `true`          | body paraments | email do usuario                                 |
| `secret`       | `string`   | `255` | `true`          | body paraments | senha do usuario                                 |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 2. /user/adm/store

criando usuário admin. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `name`         | `string`   | `255` | `true`          | body paraments | nome do usuario                                  |
| `telephone`    | `string`   | `13`  | `true`          | body paraments | telefone do usuario                              |
| `nick`         | `string`   | `255` | `true`          | body paraments | nick do usuario                                  |
| `email`        | `string`   | `255` | `true`          | body paraments | email do usuario                                 |
| `secret`       | `string`   | `255` | `true`          | body paraments | senha do usuario                                 |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 3. /user/list

listando todos os usuarios. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `count`        | `int`      | numero total de dados do banco dessa query       |
| `users`        | `[]User`   | array de users                                   |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

| User        | type     | description        |
| ----------- | -------- | ------------------ |
| `personID`  | `string` | id da pessoa       |
| `userID`    | `string` | id do usuario      |
| `name`      | `string` | nome da pessoa     |
| `telephone` | `string` | telefone da pessoa |
| `nick`      | `string` | nick do usuario    |
| `email`     | `string` | email do usuario   |
| `kind`      | `string` | tipo do usuario    |

## 4. /user/list/name/{name}

listando usuarios pelo nome. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados são trazidos         |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `count`        | `int`      | numero total de dados do banco dessa query       |
| `users`        | `[]User`   | array de users                                   |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

| User        | type     | description        |
| ----------- | -------- | ------------------ |
| `personID`  | `string` | id da pessoa       |
| `userID`    | `string` | id do usuario      |
| `name`      | `string` | nome da pessoa     |
| `telephone` | `string` | telefone da pessoa |
| `nick`      | `string` | nick do usuario    |
| `email`     | `string` | email do usuario   |
| `kind`      | `string` | tipo do usuario    |

## 5. /user/find/id/{id}

buscar um usuario pelo id. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `-`  | `true`          | url paraments     | id do usuario                                    |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `personID`     | `string`   | id da pessoa                                     |
| `userID`       | `string`   | id do usuario                                    |
| `name`         | `string`   | nome da pessoa                                   |
| `telephone`    | `string`   | telefone da pessoa                               |
| `nick`         | `string`   | nick do usuario                                  |
| `email`        | `string`   | email do usuario                                 |
| `kind`         | `string`   | tipo do usuario                                  |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 6. /user/update/id/{id}

atualizar um usuario por id.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | PUT    | yes               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `id`           | `string`   | `36`  | `true`          | url paraments  | id do usuario                                    |
| `name`         | `string`   | `255` | `true`          | body paraments | nome do usuario                                  |
| `telephone`    | `string`   | `13`  | `true`          | body paraments | telefone do usuario                              |
| `nick`         | `string`   | `255` | `true`          | body paraments | nick do usuario                                  |
| `email`        | `string`   | `255` | `true`          | body paraments | email do usuario                                 |
| `kind`         | `string`   | `10`  | `true`          | body paraments | kind do usuario                                  |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 7. /user/remove/id/{id}

remover um usuario por id.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | DELETE | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id do usuario                                    |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 8. /user/login

fazer login no sistema.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | not               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `nick`         | `string`   | `255` | `true`          | body paraments | email ou nick do usuario                         |
| `password`     | `string`   | `255` | `true`          | body paraments | senha do usuario                                 |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `token`        | `string`   | token de acesso a aplicação                      |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 9. /user/logout

fazer logout no sistema.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | true              |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 10. /user/recor/email

primeiro estagio de recuperação de senha.
enviar um codigo valido ao email do usuario.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | not               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `email`        | `string`   | `255` | `true`          | body paraments | email do usuario                                 |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 11. /user/verific/code

segundo estagio de recuperação de senha.
conferi o codigo enviado e devover um token especial.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | not               |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `code`         | `string`   | `6`  | `true`          | body paraments | codigo                                           |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `token`        | `string`   | token especial para recuperção de senha          |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 12. /user/password/recovery

terceiro estagio de recuperação de senha.
atualizar senha do usuario.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `newPassword`  | `string`   | `255` | `true`          | body paraments | nova senha do usuario                            |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 13. /user/password/update

atualizar senha do usuario.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `oldPassword`  | `string`   | `255` | `true`          | body paraments | antiga senha do usuario                          |
| `newPassword`  | `string`   | `255` | `true`          | body paraments | nova senha do usuario                            |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |


<hr>
<h1> POST Routes </h1>

## 14. /post/store

criando uma postagem. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size   | is it required? | type send      | description                                      |
| -------------- | ---------- | ------ | --------------- | -------------- | ------------------------------------------------ |
| `title`        | `string`   | `255`  | `true`          | body paraments | titulo da postagem                               |
| `content`      | `string`   | `9999` | `true`          | body paraments | conteudo da postagem                             |
| `mid`          | `string`   | `-`    | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 15. /post/list

listando todas as postagens.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | not               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `count`        | `int`      | numero total de dados do banco dessa query       |
| `posts`        | `[]Post`   | array de posts                                   |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

| Post      | type     | description                 |
| --------- | -------- | --------------------------- |
| `postID`  | `string` | id da postagem              |
| `title`   | `string` | titulo da postagem          |
| `content` | `string` | conteudo da postagem        |
| `likes`   | `int`    | numero de likes da postagem |

## 16. /post/list/title/{title}

listando todas as postagens por titulo.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | not               |

| attribute name | type value | size  | is it required? | type send         | description                                         |
| -------------- | ---------- | ----- | --------------- | ----------------- | --------------------------------------------------- |
| `title`        | `string`   | `255` | `true`          | url paraments     | titulo do post                                      |
| `offset`       | `int`      | `-`   | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`   | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`   | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`   | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `count`        | `int`      | numero total de dados do banco dessa query       |
| `posts`        | `[]Post`   | array de posts                                   |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

| Post      | type     | description                 |
| --------- | -------- | --------------------------- |
| `postID`  | `string` | id da postagem              |
| `title`   | `string` | titulo da postagem          |
| `content` | `string` | conteudo da postagem        |
| `likes`   | `int`    | numero de likes da postagem |

## 17. /post/list/category/name/{category}

listando todas as postagens por categoria.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | not               |

| attribute name | type value | size  | is it required? | type send         | description                                         |
| -------------- | ---------- | ----- | --------------- | ----------------- | --------------------------------------------------- |
| `category`     | `string`   | `255` | `true`          | url paraments     | categoria vinculada ao post                         |
| `offset`       | `int`      | `-`   | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`   | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`   | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`   | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `count`        | `int`      | numero total de dados do banco dessa query       |
| `posts`        | `[]Post`   | array de posts                                   |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

| Post      | type     | description                 |
| --------- | -------- | --------------------------- |
| `postID`  | `string` | id da postagem              |
| `title`   | `string` | titulo da postagem          |
| `content` | `string` | conteudo da postagem        |
| `likes`   | `int`    | numero de likes da postagem |

## 18. /post/find/id/{id}

buscando uma postagem pelo id. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id do post                                       |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `postID`       | `string` | id da postagem                                   |
| `title`        | `string` | titulo da postagem                               |
| `content`      | `string` | conteudo da postagem                             |
| `likes`        | `int`    | numero de likes da postagem                      |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |


## 19. /post/update/id/{id}

atualizando uma postagem. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | PUT    | yes               |

| attribute name | type value | size   | is it required? | type send      | description                                      |
| -------------- | ---------- | ------ | --------------- | -------------- | ------------------------------------------------ |
| `id`           | `string`   | `36`   | `true`          | url paraments  | id da postagem                                   |
| `title`        | `string`   | `255`  | `true`          | body paraments | titulo da postagem                               |
| `content`      | `string`   | `9999` | `true`          | body paraments | conteudo da postagem                             |
| `mid`          | `string`   | `-`    | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 20. /post/remove/id/{id}

removendo uma postagem pelo id. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | DELETE | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id do post                                       |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

<hr>
<h1> Category Routes </h1>

## 22. /category/store

criando uma categoria. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `name`         | `string`   | `255` | `true`          | body paraments | nome da categoria                                |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 22. /category/list

listando todas as categorias.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | not               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value   | description                                      |
| -------------- | ------------ | ------------------------------------------------ |
| `count`        | `int`        | numero total de dados do banco dessa query       |
| `categorys`    | `[]Category` | array de categorys                               |
| `mid`          | `string`     | mensagem da resposta caso o codigo http seja 200 |

| Category     | type     | description      |
| ------------ | -------- | ---------------- |
| `categoryID` | `string` | id da category   |
| `name`       | `string` | nome da category |

## 23. /category/find/id/{id}

buscando uma categoria pelo id. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id do categoria                                  |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `categoryID`   | `string` | id da categoria                                  |
| `name`         | `string` | nome da categoria                                |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |


## 24. /category/update/id/{id}

atualizando uma categoria. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | PUT    | yes               |

| attribute name | type value | size  | is it required? | type send      | description                                      |
| -------------- | ---------- | ----- | --------------- | -------------- | ------------------------------------------------ |
| `id`           | `string`   | `36`  | `true`          | url paraments  | id da categoria                                  |
| `name`         | `string`   | `255` | `true`          | body paraments | nome da categoria                                |
| `mid`          | `string`   | `-`   | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 25. /category/remove/id/{id}

removendo uma categoria pelo id. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | DELETE | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id da categoria                                  |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

<hr>
<h1> Post Category Routes </h1>

## 26. /post/category/store

criando uma relação entre categoria e postagem. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `postID`       | `string`   | `36` | `true`          | body paraments | id da postagem                                   |
| `categoryId`   | `string`   | `36` | `true`          | body paraments | id da categoria                                  |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 27. /category/remove/id/{id}

removendo uma categoria pelo id. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | DELETE | yes               |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `postID`       | `string`   | `36` | `true`          | body paraments | id da postagem                                   |
| `categoryId`   | `string`   | `36` | `true`          | body paraments | id da categoria                                  |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

<hr>
<h1> numberLikes Routes </h1>

## 28. /user/post/like

curtir postagem.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `postID`       | `string`   | `36` | `true`          | body paraments | id da postagem                                   |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 29. /user/post/dislike

descurtir uma postagem.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | DELETE | yes               |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `postID`       | `string`   | `36` | `true`          | body paraments | id da postagem                                   |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

<hr>
<h1> comment Routes </h1>

## 30. /comment/store

criar comentario de uma postagem.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size   | is it required? | type send      | description                                      |
| -------------- | ---------- | ------ | --------------- | -------------- | ------------------------------------------------ |
| `postID`       | `string`   | `36`   | `true`          | body paraments | id da postagem                                   |
| `title`        | `string`   | `255`  | `true`          | body paraments | titulo do comentario                             |
| `content`      | `string`   | `2024` | `true`          | body paraments | conteudo do comentario                           |
| `mid`          | `string`   | `-`    | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 31. /comment/list/post/id/{postID}

listando todos os comentarios de uma postagem.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | not               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `postID`       | `string`   | `36` | `true`          | url paraments     | id da postagem                                      |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value  | description                                      |
| -------------- | ----------- | ------------------------------------------------ |
| `count`        | `int`       | numero total de dados do banco dessa query       |
| `comments`     | `[]Comment` | array de comments                                |
| `mid`          | `string`    | mensagem da resposta caso o codigo http seja 200 |

| Comment     | type     | description          |
| ----------- | -------- | -------------------- |
| `commentID` | `string` | id do comentario     |
| `title`     | `string` | titulo da postagem   |
| `content`   | `string` | conteudo da postagem |
| `userID`    | `string` | id do usuario        |
| `postID`    | `string` | id do post           |

## 32. /comment/list/user

listando todos os comentarios do usuario.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value  | description                                      |
| -------------- | ----------- | ------------------------------------------------ |
| `count`        | `int`       | numero total de dados do banco dessa query       |
| `comments`     | `[]Comment` | array de comments                                |
| `mid`          | `string`    | mensagem da resposta caso o codigo http seja 200 |

| Comment     | type     | description          |
| ----------- | -------- | -------------------- |
| `commentID` | `string` | id do comentario     |
| `title`     | `string` | titulo da postagem   |
| `content`   | `string` | conteudo da postagem |
| `userID`    | `string` | id do usuario        |
| `postID`    | `string` | id do post           |

## 33. /comment/list/user/post/id/{postID}

listando todos os comentarios de um usuario em uma postagem.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `postID`       | `string`   | `36` | `true`          | url paraments     | id do post                                          |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value  | description                                      |
| -------------- | ----------- | ------------------------------------------------ |
| `count`        | `int`       | numero total de dados do banco dessa query       |
| `comments`     | `[]Comment` | array de comments                                |
| `mid`          | `string`    | mensagem da resposta caso o codigo http seja 200 |

| Comment     | type     | description          |
| ----------- | -------- | -------------------- |
| `commentID` | `string` | id do comentario     |
| `title`     | `string` | titulo da postagem   |
| `content`   | `string` | conteudo da postagem |
| `userID`    | `string` | id do usuario        |
| `postID`    | `string` | id do post           |

## 34. /comment/find/id/{id}

buscando um comentario pelo id.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id do post                                       |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `commentID`    | `string` | id do comentario                                 |
| `title`        | `string` | titulo da postagem                               |
| `content`      | `string` | conteudo da postagem                             |
| `userID`       | `string` | id do usuario                                    |
| `postID`       | `string` | id do post                                       |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

## 35. /comment/update/id/{id}

atualizando um comentario.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | PUT    | yes               |

| attribute name | type value | size   | is it required? | type send      | description                                      |
| -------------- | ---------- | ------ | --------------- | -------------- | ------------------------------------------------ |
| `id`           | `string`   | `36`   | `true`          | url paraments  | id da categoria                                  |
| `title`        | `string`   | `255`  | `true`          | body paraments | titulo do comentario                             |
| `content`      | `string`   | `2024` | `true`          | body paraments | conteudo do comentario                           |
| `mid`          | `string`   | `-`    | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 36. /comment/remove/id/{id}

remover um comentario.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | DELETE | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id do comentario                                 |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

<hr>
<h1> response comment Routes </h1>

## 37. /response/comment/store

responder a um comentario de um usuario.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size   | is it required? | type send      | description                                      |
| -------------- | ---------- | ------ | --------------- | -------------- | ------------------------------------------------ |
| `commentId`    | `string`   | `36`   | `true`          | body paraments | id do comentario                                 |
| `title`        | `string`   | `255`  | `true`          | body paraments | titulo do comentario de resposta                 |
| `content`      | `string`   | `2024` | `true`          | body paraments | conteudo do comentario  de resposta              |
| `mid`          | `string`   | `-`    | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 38. /response/comment/list/comment/id/{commentID}

listando todos os comentarios de resposta de um comentario.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | not               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `commentID`    | `string`   | `36` | `true`          | url paraments     | id do comentario                                    |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name     | type value          | description                                      |
| ------------------ | ------------------- | ------------------------------------------------ |
| `count`            | `int`               | numero total de dados do banco dessa query       |
| `responsecomments` | `[]ResponseComment` | array de response comments                       |
| `mid`              | `string`            | mensagem da resposta caso o codigo http seja 200 |

| ResponseComment     | type     | description                        |
| ------------------- | -------- | ---------------------------------- |
| `responseCommentID` | `string` | id do comentario de resposta       |
| `title`             | `string` | titulo do comentario de resposta   |
| `content`           | `string` | conteudo do comentario de resposta |
| `commentID`         | `string` | id do comentario                   |
| `userID`            | `string` | id do usuario                      |

## 39. /response/comment/list/user

listando todos os comentarios de resposta de um usuario.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name     | type value          | description                                      |
| ------------------ | ------------------- | ------------------------------------------------ |
| `count`            | `int`               | numero total de dados do banco dessa query       |
| `responsecomments` | `[]ResponseComment` | array de response comments                       |
| `mid`              | `string`            | mensagem da resposta caso o codigo http seja 200 |

| ResponseComment     | type     | description                        |
| ------------------- | -------- | ---------------------------------- |
| `responseCommentID` | `string` | id do comentario de resposta       |
| `title`             | `string` | titulo do comentario de resposta   |
| `content`           | `string` | conteudo do comentario de resposta |
| `commentID`         | `string` | id do comentario                   |
| `userID`            | `string` | id do usuario                      |


## 40. /response/comment/update/id/{id}

atualizando um comentario de resposta.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | PUT    | yes               |

| attribute name | type value | size   | is it required? | type send      | description                                      |
| -------------- | ---------- | ------ | --------------- | -------------- | ------------------------------------------------ |
| `id`           | `string`   | `36`   | `true`          | url paraments  | id do comentario de resposta                     |
| `title`        | `string`   | `255`  | `true`          | body paraments | titulo do comentario de resposta                 |
| `content`      | `string`   | `2024` | `true`          | body paraments | conteudo do comentario de resposta               |
| `mid`          | `string`   | `-`    | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 41. /response/comment/remove/id/{id}

remover um comentario de resposta.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | DELETE | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id do comentario de resposta                     |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

<hr>
<h1> config Routes </h1>

## 42. /config/store

criar uma configuração. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | POST   | yes               |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `collors`      | `[]string` | `-`  | `true`          | body paraments | cores do site                                    |
| `links`        | `[]string` | `-`  | `true`          | body paraments | links do menu do site                            |
| `menuAncoras`  | `[]string` | `-`  | `true`          | body paraments | ancoras do menu                                  |
| `banner`       | `string`   | `-`  | `true`          | body paraments | url do banner                                    |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 43. /config/list

listando todas as configurações. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                         |
| -------------- | ---------- | ---- | --------------- | ----------------- | --------------------------------------------------- |
| `offset`       | `int`      | `-`  | `false`         | queries paraments | deslocamento inicial dos dados trazidos             |
| `limit`        | `int`      | `-`  | `false`         | queries paraments | limite padrão de quantos dados serão trazidos       |
| `page`         | `int`      | `-`  | `false`         | queries paraments | o numero da pagina na qual os dados estão agrupados |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200    |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `count`        | `int`      | numero total de dados do banco dessa query       |
| `configs`      | `[]Config` | array de configs                                 |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

| Config        | type     | description        |
| ------------- | -------- | ------------------ |
| `configID`    | `string` | id da configuração |
| `collors`     | `string` | cores do site      |
| `links`       | `string` | links do menu      |
| `menuAncoras` | `string` | ancoras do menu    |
| `banner`      | `string` | url do banner      |

## 44. /config/find/id/{id}

buscando uma configuração pelo id. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| queries | -    | GET    | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `int`      | `-`  | `true`          | queries paraments | id da configuração                               |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `configID`     | `string` | id da configuração                               |
| `collors`      | `string` | cores do site                                    |
| `links`        | `string` | links do menu                                    |
| `menuAncoras`  | `string` | ancoras do menu                                  |
| `banner`       | `string` | url do banner                                    |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |


## 45. /config/update/id/{id}

atualizando uma configuração. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type   | method | token is required |
| ------- | ------ | ------ | ----------------- |
| body    | object | PUT    | yes               |

| attribute name | type value | size | is it required? | type send      | description                                      |
| -------------- | ---------- | ---- | --------------- | -------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments  | id da configuração                               |
| `collors`      | `[]string` | `-`  | `true`          | body paraments | cores do site                                    |
| `links`        | `[]string` | `-`  | `true`          | body paraments | links do menu do site                            |
| `menuAncoras`  | `[]string` | `-`  | `true`          | body paraments | ancoras do menu                                  |
| `banner`       | `string`   | `-`  | `true`          | body paraments | url do banner                                    |
| `mid`          | `string`   | `-`  | `false`         | body paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |

| attribute name | type value | description                                      |
| -------------- | ---------- | ------------------------------------------------ |
| `mid`          | `string`   | mensagem da resposta caso o codigo http seja 200 |

## 46. /config/remove/id/{id}

remover uma configuração. <br>
somente usuario admin pode utilizar esse endpoint.

#### - _Request_

| request | type | method | token is required |
| ------- | ---- | ------ | ----------------- |
| body    | -    | DELETE | yes               |

| attribute name | type value | size | is it required? | type send         | description                                      |
| -------------- | ---------- | ---- | --------------- | ----------------- | ------------------------------------------------ |
| `id`           | `string`   | `36` | `true`          | url paraments     | id da configuração                               |
| `mid`          | `string`   | `-`  | `false`         | queries paraments | mensagem da resposta caso o codigo http seja 200 |

#### - _Response_

| request | type   | status |
| ------- | ------ | ------ |
| body    | object | 200    |


| attribute name | type     | description                                      |
| -------------- | -------- | ------------------------------------------------ |
| `mid`          | `string` | mensagem da resposta caso o codigo http seja 200 |

the end!
made by Jonatas.
