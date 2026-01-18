1️⃣ Struct + Método (básico)

Crie uma struct chamada Pessoa com:

Nome string

Idade int

Implemente um método:

func (p Pessoa) EhMaiorDeIdade() bool


📌 Deve retornar true se idade ≥ 18.

2️⃣ Método com ponteiro

Crie uma struct ContaBancaria com:

Saldo float64

Implemente um método com receiver ponteiro:

func (c *ContaBancaria) Depositar(valor float64)


📌 O saldo deve ser alterado corretamente.

3️⃣ Aninhamento de structs

Crie:

type Endereco struct {
    Cidade string
    Estado string
}

type Cliente struct {
    Nome string
    Endereco
}


📌 Acesse Cidade sem usar cliente.Endereco.Cidade.

4️⃣ Interface simples

Crie uma interface:

type Animal interface {
    Falar() string
}


Implemente a interface para:

Cachorro

Gato

📌 Cada um retorna um som diferente.

5️⃣ Polimorfismo real

Crie uma função:

func EmitirSom(a Animal)


📌 Ela deve imprimir o resultado de a.Falar()
Use com tipos diferentes.

6️⃣ Slice de interface

Crie um slice:

var animais []Animal


📌 Adicione pelo menos dois tipos diferentes que implementem Animal
Percorra o slice e chame Falar().

7️⃣ Interface + Struct embutida

Crie:

type Veiculo interface {
    VelocidadeMaxima() int
}


Crie uma struct base:

type Motor struct {
    Potencia int
}


E uma struct Carro que embuta Motor e implemente Veiculo.

8️⃣ Interface aplicada (mais concreta)

Crie a interface:

type Pagamento interface {
    Pagar(valor float64) bool
}


Crie duas structs:

CartaoCredito

Pix

📌 Cada uma implementa Pagar de forma diferente (mensagens diferentes já servem).

Crie a função:

func ProcessarPagamento(p Pagamento, valor float64)


👉 Ela deve chamar p.Pagar(valor).

9️⃣ Type assertion com caso real

Crie uma função:

func ProcessarValor(v interface{})


📌 Comportamento:

se for int → imprimir Inteiro

se for float64 → imprimir Float

se não for nenhum → imprimir Tipo não suportado

👉 Use type assertion segura, não type switch ainda.

🔟 Type switch aplicado

Crie uma função:

func Log(v interface{})


📌 Use switch v.(type) para:

string → imprimir o tamanho

int → imprimir o dobro

bool → imprimir se é verdadeiro ou falso

default → imprimir Tipo desconhecido

➕ NOVOS DESAFIOS (11 → 15) – dificuldade crescente
1️⃣1️⃣ Interface + estado interno (nível médio)

Crie uma interface:

type Dispositivo interface {
    Ligar()
    Desligar()
}


Implemente:

TV

Computador

📌 Cada um deve manter um estado interno (ligado bool).

1️⃣2️⃣ Método com ponteiro + interface

Crie:

type Contador interface {
    Incrementar()
    Valor() int
}


Implemente uma struct que só funcione corretamente com ponteiro.

📌 Teste chamando via interface.

1️⃣3️⃣ Interface como retorno de função

Crie duas structs:

EmailNotificador

SMSNotificador

Implemente:

type Notificador interface {
    Enviar(msg string)
}


Crie a função:

func NovoNotificador(tipo string) Notificador


📌 Retorne implementações diferentes.

1️⃣4️⃣ Slice de interface + type assertion

Crie:

var itens []interface{}


Adicione:

int

string

struct qualquer

📌 Percorra e:

se for int → multiplique

se for string → imprima em maiúsculo

se for struct → imprima o tipo

👉 Aqui você mistura assertion + lógica.

1️⃣5️⃣ Composição + override de comportamento (nível alto)

Crie:

type Logger interface {
    Log(msg string)
}

Crie:

LoggerBase

LoggerArquivo

LoggerConsole

📌 Um deles deve reutilizar o comportamento base e alterar parte da lógica.

👉 Isso força você a entender:

composição

método com mesmo nome

chamada explícita do método embutido
