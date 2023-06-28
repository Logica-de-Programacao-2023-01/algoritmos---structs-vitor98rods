package main

import "fmt"
type MUS struct {
	Titulo  string
	Artista string
	Duracao string
}

type Playlist struct {
	Nome    string
	Musicas []MUS
}

func buscarPlaylistsPorTitulo(tituloMusica string, playlists []Playlist) []Playlist {
	resultado := []Playlist{}

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == tituloMusica {
				resultado = append(resultado, playlist)
				break
			}
		}
	}

	return resultado
}

func main() {
	musicas1 := []MUS{
		MUS{Titulo: "Música 1", Artista: "Artista 1", Duracao: "3:30"},
		MUS{Titulo: "Música 2", Artista: "Artista 2", Duracao: "4:15"},
		MUS{Titulo: "Música 3", Artista: "Artista 3", Duracao: "5:10"},
	}

	musicas2 := []MUS{
		MUS{Titulo: "Música 4", Artista: "Artista 4", Duracao: "2:45"},
		MUS{Titulo: "Música 5", Artista: "Artista 5", Duracao: "3:20"},
	}

	playlist1 := Playlist{Nome: "Playlist 1", Musicas: musicas1}
	playlist2 := Playlist{Nome: "Playlist 2", Musicas: musicas2}

	playlists := []Playlist{playlist1, playlist2}

	titulo := "Música 3"
	playlistsEncontradas := buscarPlaylistsPorTitulo(titulo, playlists)

	if len(playlistsEncontradas) > 0 {
		fmt.Printf("Playlists encontradas com a música '%s':\n", titulo)
		for _, playlist := range playlistsEncontradas {
			fmt.Println(playlist.Nome)
		}
	} else {
		fmt.Printf("Nenhuma playlist encontrada com a música '%s'\n", titulo)
	}
}
