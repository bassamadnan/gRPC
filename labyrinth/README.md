# Labyrinth
The following are the possible states of a cell:
- H (Hidden, only seen on client side)
- W (Wall)
- C (Coin)
- P (Player position)
- G (Goal)

Make sure to run `go mod tidy`
## Server
```bash
go run cmd/server/main.go
```
Start by running the server, which taeks input from the `grid.txt` , the original grid can be seen on server terminal.

After spells are cast (Bombarda) the cell is printed again (for debugging purposes) client activity can also be seen

## Client
```bash
go run cmd/server/client.go
```
After running the server, start the client. The client can be moved by WASD keys.
### Spell casting
- Revelio : Press 1 and enter the cell position and the type of cell (eg: `1 2 C`) space seperated
- Bombarda : Press 2 and enter three cells (two at a time, and press Enter), expected to be valid cells

To restart the game, you will have to restart the server since the server maintains a central state of the game board.
