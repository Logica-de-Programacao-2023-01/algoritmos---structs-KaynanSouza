package main

import (
	"fmt"
	"time"
)

// Crie uma struct chamada Playlist com os campos "nome" e "músicas".
// O campo "músicas" deve ser um slice de structs, cada uma representando
// uma música com os campos "título", "artista" e "duração". Escreva uma função
// que receba uma Playlist como parâmetro e imprima o nome da playlist, o tempo
// total da playlist e as informações de cada música.
func main() {
	playlist := playlist{
		nome: "teste",
		musica: []musica{
			{titulo: "titulo1", artista: "artista1", duracao: 1},
			{titulo: "titulo2", artista: "artista2", duracao: 2},
			{titulo: "titulo3", artista: "artista3", duracao: 3},
		},
	}
	PrintPlaylist(playlist)
}
func (p playlist) DuracaoT() time.Duration {
	total := time.Duration(0)
	for _, r := range p.musica {
		total += r.duracao
	}
	return total
}
func PrintPlaylist(p playlist) {
	fmt.Println(p.nome)
	fmt.Println(p.DuracaoT)
	for _, m := range p.musica {
		fmt.Println(m.titulo, m.titulo, m.duracao)
	}
}

type playlist struct {
	nome   string
	musica []musica
}
type musica struct {
	titulo  string
	artista string
	duracao time.Duration
}
