package gamestate

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/ken343/maritime-royale/pkg/elements"
	"github.com/ken343/maritime-royale/pkg/networking/mrp"
)

var ObjectMap = make(map[string]*elements.Element)

type Chunk struct {
	ChunkID          string
	ChunkTerrainData []*elements.Element
	ChunkUnitData    []*elements.Element
}

var chunkList []*Chunk
var chunkListTemp []*Chunk

func GetEntireWorld() []*elements.Element {
	var masterMap = []*elements.Element{}

	for _, chunk := range chunkList {

		masterMap = append(masterMap, chunk.ChunkTerrainData...)
		masterMap = append(masterMap, chunk.ChunkUnitData...)
	}

	return masterMap
}

func CreateChunk() {
	ID := strconv.Itoa(len(chunkListTemp))
	chunkListTemp = append(chunkListTemp, &Chunk{ChunkID: ID})
	chunkList = append(chunkList, &Chunk{ChunkID: ID})
}

func AddUnitToWorld(elem *elements.Element) {

	if len(chunkListTemp) == 0 {
		CreateChunk()
	}
	mu.Lock()
	for _, name := range blacklistedNames {
		if elem.UniqueName == name {
			mu.Unlock()
			return
		}
	}
	mu.Unlock()
	for _, unitElem := range chunkList[0].ChunkUnitData {
		if unitElem.UniqueName == elem.UniqueName {
			mu.Lock()
			*unitElem = *elem
			mu.Unlock()
			return
		}
	}

	mu.Lock()
	chunkList[0].ChunkUnitData = append(chunkList[0].ChunkUnitData, elem)
	mu.Unlock()

}

func AddTerrainToWorld(elem *elements.Element) {
	if len(chunkListTemp) == 0 {
		CreateChunk()
	}
	mu.Lock()
	for _, name := range blacklistedNames {
		if elem.UniqueName == name {
			mu.Unlock()
			return
		}
	}
	mu.Unlock()

	for _, unitElem := range chunkList[0].ChunkTerrainData {
		if unitElem.UniqueName == elem.UniqueName {
			mu.Lock()
			*unitElem = *elem
			mu.Unlock()
			return
		}
	}

	mu.Lock()
	chunkList[0].ChunkTerrainData = append(chunkList[0].ChunkTerrainData, elem)
	mu.Unlock()
}

func GetObject(objectName string) *elements.Element {
	var returnElem *elements.Element = new(elements.Element)
	returnElem = ObjectMap[objectName].MakeCopy()
	return returnElem
}

func RemoveElem(badElem *elements.Element) {
	bytes, _ := json.Marshal(&badElem)

	myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte("NIL"))

	for k, existing := range chunkList[0].ChunkUnitData {
		if badElem.UniqueName == existing.UniqueName {
			if k < len(chunkList[0].ChunkUnitData) {
				copy(chunkList[0].ChunkUnitData[k:], chunkList[0].ChunkUnitData[k+1:])
			}
			chunkList[0].ChunkUnitData[len(chunkList[0].ChunkUnitData)-1] = nil
			chunkList[0].ChunkUnitData = chunkList[0].ChunkUnitData[:len(chunkList[0].ChunkUnitData)-1]

		}
	}

	for _, conn := range connectionList {
		conn.Write(myMRP.MRPToByte())
	}
}

var blacklistedNames []string

var mu sync.Mutex

func PushChunks() {
	mu.Lock()
	mu.Unlock()
}
