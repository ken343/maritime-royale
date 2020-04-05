package gamestate

import (
	"encoding/json"
	"log"
	"net"
	"strconv"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/networking/mrp"
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
	chunkListTemp[0].ChunkUnitData = append(chunkListTemp[0].ChunkUnitData, elem)
}

func AddTerrainToWorld(elem *elements.Element) {
	if len(chunkListTemp) == 0 {
		CreateChunk()
	}
	chunkListTemp[0].ChunkTerrainData = append(chunkListTemp[0].ChunkTerrainData, elem)
}

func GetObject(objectName string) *elements.Element {
	var returnElem *elements.Element = new(elements.Element)
	returnElem = ObjectMap[objectName].MakeCopy()
	return returnElem
}

func PushChunks() {
	var found bool = false
	for _, chunkTemp := range chunkListTemp {
		for _, chunk := range chunkList {
			if chunkTemp.ChunkID == chunk.ChunkID {
				//sync Terrain
				for _, TerrainElemTemp := range chunkTemp.ChunkTerrainData {
					for _, TerrainElem := range chunk.ChunkTerrainData {
						if TerrainElem.UniqueName == TerrainElemTemp.UniqueName {
							*TerrainElem = *TerrainElemTemp
							found = true
							break
						}
					}
					if found == true {
						found = false
					} else {
						chunk.ChunkTerrainData = append(chunk.ChunkTerrainData, TerrainElemTemp)
					}
				}
				//sync Units
				for _, unitElemTemp := range chunkTemp.ChunkUnitData {
					for _, unitElem := range chunk.ChunkUnitData {
						if unitElem.UniqueName == unitElemTemp.UniqueName {
							*unitElem = *unitElemTemp
							found = true
							break
						}
					}
					if found == true {
						found = false
					} else {
						chunk.ChunkUnitData = append(chunk.ChunkUnitData, unitElemTemp)
					}
				}
			}
		}
		chunkTemp.ChunkUnitData = []*elements.Element{}
		chunkTemp.ChunkTerrainData = []*elements.Element{}
	}
}

func SendElemMap(conn net.Conn) {
	myMap := GetEntireWorld()

	for _, myElem := range myMap {
		bytes, err := json.Marshal(myElem)
		if err != nil {
			log.Fatal(err)
		}

		myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(""))
		conn.Write(myMRP.MRPToByte())

	}

	ForceUpdate(conn)
}
