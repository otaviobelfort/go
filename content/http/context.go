package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	/*ctx, cancel := ContexHotel2()
	defer cancel()
	BookHotel(ctx)
	ContexHotel()*/
	ServerContext()

}
func ServerContext() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("Request is processing...")
	defer fmt.Println("Request is processed!")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Requisição processada com sucesso!")
		w.Write([]byte("Requisição processada com sucesso!"))
	case <-ctx.Done():
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}

}

func ContexHotel() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	BookHotel(ctx)
}
func ContexHotel2() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	return ctx, cancel

}

func BookHotel(ctx context.Context) {
	select {
	// quando passar o timeout de 3 segundos
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("Hotel booked successfully")
	}

}
