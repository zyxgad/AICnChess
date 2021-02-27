
package chess

import (
	bytes "bytes"
)


type ChessBoard struct{
	pieces [10][9]byte
	movelogs []*ChessLog

	b_gen_Z byte
	r_gen_Z byte
	cycle byte
}

func InitBoard()(board *ChessBoard){
	board = new(ChessBoard)
	board.ClearBoard()
	return board
}
func InitBoardFromBytes(buf *bytes.Buffer)(board *ChessBoard){
	board = InitBoard()
	for i := 0; i < len(board.pieces) ;i++ {
		for j := 0; j < len(board.pieces[0]) ;j++ {
			piece, _ := buf.ReadByte()
			board.pieces[i][j] = piece
			if piece & PIECE_GENERAL != 0 {
				if piece & TEAM_BLK != 0 {
					board.b_gen_Z = zipByte2(&[2]byte{byte(i), byte(j)})
				}else if piece & TEAM_RED != 0 {
					board.r_gen_Z = zipByte2(&[2]byte{byte(i), byte(j)})
				}
			}
		}
	}
	board.cycle, _ = buf.ReadByte()
	loglength := int(readUint32FromBuf(buf))
	board.movelogs = make([]*ChessLog, 0, loglength)
	for i := 0; i < loglength ;i++ {
		board.movelogs = append(board.movelogs, InitChessLogFromBytes(buf))
	}
	return board
}
func (board *ChessBoard)ToBytes(buf *bytes.Buffer)(*bytes.Buffer){
	for _, row := range board.pieces {
		for _, p := range row {
			buf.WriteByte(p)
		}
	}
	buf.WriteByte(board.cycle)

	writeUint32ToBuf(buf, uint32(len(board.movelogs)))
	for _, l := range board.movelogs {
		l.ToBytes(buf)
	}
	return buf
}

func (board *ChessBoard)ClearBoard(){
	for i := 0; i < 10 ;i++ {
		for j := 0; j < 9 ;j++{
			board.pieces[i][j] = PIECE_NULL
		}
	}
	board.cycle = TEAM_NULL;
}

func (board *ChessBoard)ResetBoard(){
	for i := 0; i < 5 ;i++{
		for j := 0; j < 9 ;j++{
			board.pieces[i][j] = stdmap[i][j] | TEAM_BLK
		}
	}
	for i := 0; i < 5 ;i++{
		for j := 0; j < 9 ;j++{
			board.pieces[10 - i][9 - j] = stdmap[i][j] | TEAM_RED
		}
	}
	board.cycle = TEAM_BLK;
}

func (board *ChessBoard)InDanger() byte {
	if board.cycle == TEAM_NULL {
		return TEAM_NULL
	}

	b_gen := uzipByte2(board.b_gen_Z)
	r_gen := uzipByte2(board.r_gen_Z)

	var flag bool = false
	if b_gen[1] == r_gen[1] {
		flag = true
		x := b_gen[1]
		for y := b_gen[0]; y < r_gen[1] ;y++{
			if board.pieces[y][x] != PIECE_NULL {
				flag = false
				break
			}
		}
		if flag {
			return checkCond(board.cycle == TEAM_BLK, TEAM_BLK, TEAM_RED).(byte)
		}
	}
	return TEAM_NULL
}

func (board *ChessBoard)CanMoveList(pos_Z byte)([]byte){
	pos := uzipByte2(pos_Z)
	if (pos[0] < 0 || pos[0] >= 10) || (pos[1] < 0 || pos[1] >= 9) {
		return nil
	}
	piece := board.pieces[pos[0]][pos[1]]
	if piece & board.cycle == 0 {
		return nil
	}
	list := make([]byte, 8)
	appendPos := func(tpos *[2]byte){
		if (board.pieces[tpos[0]][tpos[1]] & board.cycle == 0) && (board.InDanger() != board.cycle){
			list = append(list, zipByte2(tpos))
		}
	}
	switch piece {
	case PIECE_SOLDIER:   // 卒/兵
		if board.cycle == TEAM_BLK {
			appendPos(&[2]byte{pos[0] + 1, pos[1]})
			if pos[0] > 5 {
				appendPos(&[2]byte{pos[0], pos[1] - 1})
				appendPos(&[2]byte{pos[0], pos[1] + 1})
			}
		}else{
			appendPos(&[2]byte{pos[0] - 1, pos[1]})
			if pos[0] <= 5 {
				appendPos(&[2]byte{pos[0], pos[1] - 1})
				appendPos(&[2]byte{pos[0], pos[1] + 1})
			}
		}
	case PIECE_CANNON:    // 炮
		//
	case PIECE_ROOK:      // 车
	case PIECE_SOWAR:     // 马
	case PIECE_MINISTER:  // 象/相
	case PIECE_GUARD:     // 士
	case PIECE_GENERAL:   // 将/帅
		if pos[0] > 3 {
			appendPos(&[2]byte{pos[0] - 1, pos[1]})
		}
		if pos[0] < 5 {
			appendPos(&[2]byte{pos[0] + 1, pos[1]})
		}
		if board.cycle == TEAM_BLK {
			if pos[1] > 0 {
				appendPos(&[2]byte{pos[0], pos[1] - 1})
			}
			if pos[1] < 2 {
				appendPos(&[2]byte{pos[0], pos[1] + 1})
			}
		}else{
			if pos[1] > 7 {
				appendPos(&[2]byte{pos[0], pos[1] - 1})
			}
			if pos[1] < 9 {
				appendPos(&[2]byte{pos[0], pos[1] + 1})
			}
		}
	default:
		return nil
	}
	return list
}

func (board *ChessBoard)MovePiece(start_Z byte, end_Z byte) bool {
	starts := uzipByte2(start_Z)
	ends   := uzipByte2(end_Z)
	if (((starts[0] < 0 || starts[0] >= 10) || (starts[1] < 0 || starts[1] >= 9)) ||
		((ends[0]   < 0 || ends[0]   >= 10) || (ends[1]   < 0 || ends[1]   >= 9))) {
		return false
	}
	end_piece := board.pieces[ends[0]][ends[1]]
	if end_piece & board.cycle != 0 {
		return false
	}
	start_piece := board.pieces[starts[0]][starts[1]]
	if (start_piece == PIECE_NULL) || (start_piece & board.cycle == 0) {
		return false
	}
	board.movelogs = append(board.movelogs, InitChessLog(board.cycle, start_Z, end_Z, start_piece, end_piece))
	return true
}
