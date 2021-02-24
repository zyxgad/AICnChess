
package chess

const(
	PIECE_NULL = iota // 空
	PIECE_SOLDIER     // 卒/兵
	PIECE_CANNON      // 炮
	PIECE_ROOK        // 车
	PIECE_SOWAR       // 马
	PIECE_MINISTER    // 象/相
	PIECE_GUARD       // 士
	PIECE_GENERAL     // 将/帅
)

type ChessMap struct{
	data [9][10]byte
}

func (cmap *ChessMap)ClearMap(){
	for i := 0; i < 10 ;i++ {
		for j := 0; j < 8 ;j++{
			cmap.data[i][j] = PIECE_NULL
		}
	}
}
func (cmap *ChessMap)ResetMap(){
	cmap.ClearMap()
}

