
package chess

import (
	bytes "bytes"
)

type ChessLog struct{
	team  byte
	start byte
	end   byte
	piece byte
	klpie byte
}

func InitChessLog(team byte, start byte, end byte, piece byte, klpie byte)(clog *ChessLog){
	clog = new(ChessLog)
	if (team == TEAM_NULL ||
		piece == PIECE_NULL) {
		return nil
	}
	clog.team   = team
	clog.start  = start
	clog.end    = end
	clog.piece  = piece
	clog.klpie  = klpie
	return clog
}

func InitChessLogFromBytes(buf *bytes.Buffer)(*ChessLog){
	team , _ := buf.ReadByte()
	start, _ := buf.ReadByte()
	end  , _ := buf.ReadByte()
	piece, _ := buf.ReadByte()
	klpie, _ := buf.ReadByte()
	return InitChessLog(team, start, end, piece, klpie)
}
func (clog *ChessLog)ToBytes(buf *bytes.Buffer)(*bytes.Buffer){
	buf.WriteByte(clog.team)
	buf.WriteByte(clog.start)
	buf.WriteByte(clog.end)
	buf.WriteByte(clog.piece)
	buf.WriteByte(clog.klpie)
	return buf
}


func (clog *ChessLog)ToStr()(str string){
	if clog.team == TEAM_BLK {
		str += "BLK "
	}else if clog.team == TEAM_RED {
		str += "RED "
	}else{
		return ""
	}
	switch clog.piece {
	case PIECE_SOLDIER:   // 卒/兵
		str += "soldier "
	case PIECE_CANNON:    // 炮
		str += "cannon "
	case PIECE_ROOK:      // 车
		str += "rook "
	case PIECE_SOWAR:     // 马
		str += "sowar "
	case PIECE_MINISTER:  // 象/相
		str += "minister "
	case PIECE_GUARD:     // 士/仕
		str += "guard "
	case PIECE_GENERAL:   // 将/帅
		str += "generl "
	default:
		return ""
	}
	start_ := uzipByte2(clog.start)
	end_   := uzipByte2(clog.end)
	str += (string('A' + start_[0]) + string('1' + start_[1]) +
			" moved to " +
			string('A' + end_[0])   + string('1' + end_[1]) )
	switch clog.klpie {
	case PIECE_SOLDIER:   // 卒/兵
		str += " kill blk soldier"
	case PIECE_CANNON:    // 炮
		str += " kill blk cannon"
	case PIECE_ROOK:      // 车
		str += " kill blk rook"
	case PIECE_SOWAR:     // 马
		str += " kill blk sowar"
	case PIECE_MINISTER:  // 象/相
		str += " kill blk minister"
	case PIECE_GUARD:     // 士/仕
		str += " kill blk guard"
	case PIECE_GENERAL:   // 将/帅
		str += " kill blk generl"
	}
	return str
}
