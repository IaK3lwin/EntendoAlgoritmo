## Caso base e caso recursivo

<p>Por conta de ser recursivo, é normal acabar escrevendo uma função recursiva
erroniamente. Causando um Loop infinito! Imagine que temos um <a href="./contagem_regressiva_errado.py">
    algoritmo que faça
    uma contagem regressiva
</a>: 3... 2... 1... 0 </p>

``` lang="python"
def regressiva(i)
    print(i)
    regressiva(i - 1)
``` 

<p>
    Nesse caso temos uma função recursiva codificada de forma errada!
    Por isso funções recursivas tem duas partes! O caso <strong>Base</strong> e o caso <strong>recursivo</strong>
</p>

## Caso Base
<p>O caso base é a parte da função recursiva que não se chama, assim acabando com o loop!</p>

``` lang="python"
    def Recursao(i):
        if i < 0:
            return // para a recursão aqui!
        ....
``` 
## Caso recursivo

<p>Parte que chama a própria função, fazendo com que a recursão aconteça de fato!</p>

```
    ....
    Recurssao(i - 1)

Recursao(3)
``` 

