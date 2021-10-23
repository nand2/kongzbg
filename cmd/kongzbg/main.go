package main

import (
	"log"
	"fmt"
	"os"
	"image"
	"encoding/json"
	"io/ioutil"
	// "image/gif"
	"image/png"
	"image/color"
)

const kongzPixelInRealPixels = 20




func main() {
	log.Printf("Hello!")

	processKongz(65)
}

func processKongz(id uint) {
	log.Printf("Processing kongz %d", id)

    encodedImageFile, err := os.Open(fmt.Sprintf("assets/images/%d", id))
    if err != nil {
        log.Panic(err)
    }
    defer encodedImageFile.Close()

    imageData, _, err := image.Decode(encodedImageFile)
    if err != nil {
        log.Panic(err)
    }

    // Copy the image in an editable data struct
	newImageData := image.NewRGBA(imageData.Bounds())
	size := newImageData.Bounds().Size()
	for x := 0; x < size.X; x++ {
	    for y := 0; y < size.Y; y++ {
	        newImageData.Set(x, y, imageData.At(x, y))
	    }
	}

	// Set the background as transparent
    for kongzY, maskLine := range getKongzBackgroundMask(id) {
    	for kongzX, mask := range maskLine {
    		if mask == 1 {
				for x := kongzX * kongzPixelInRealPixels; x < (kongzX + 1) * kongzPixelInRealPixels; x++ {
				    for y := kongzY * kongzPixelInRealPixels; y < (kongzY + 1) * kongzPixelInRealPixels; y++ {
				        newImageData.Set(x, y, color.RGBA{0, 0, 0, 0})
				    }
				}    			
    		}
    	}
    }

    // Save the result
    outfile, err := os.Create(fmt.Sprintf("assets/images-nobg/%d", id))
    if err != nil {
        log.Panic(err)
    }
    defer outfile.Close()
    png.Encode(outfile, newImageData)
}

func getKongzBackgroundMask(id uint) (mask [][]int) {


	// Common mask to all
	var mainBackgroundMask = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
	}

	// Init the mask
	mask = mainBackgroundMask

	// Get the genesis traits
	kongzTraits := getKongzTraitByType("Genesis Trait", id)
	log.Printf("Traits: %v", kongzTraits)

	// Substractive mask per trait
	var traitsSubstractiveMasks = []struct{
		traits []string
		maskLocationsToUnset [][]int // X/Y format
	}{
		{
			traits: []string{"Propeller Hat"},
			maskLocationsToUnset: [][]int{{5, 10}, {5, 9}, {5, 8}, {6, 8}, {6, 7}, {7, 7}, {7, 6}, {8, 7}, {8, 6}, {8, 5}, {9, 6}, {9, 5}, {9, 4}, {10, 6}, {10, 5}, {10, 4}, {11, 5}, {11, 4}, {11, 3}, {12, 4}, {12, 3}, {13, 4}, {13, 3}, {14, 4}, {14, 3}, {15, 4}, {15, 3}, {16, 3}, {17, 3}, {18, 3}, {19, 3}, {20, 3}, {21, 3}, {22, 3}, {23, 4}, {24, 4}, {25, 5}, {26, 7}, {26, 6}, {27, 8}, {27, 7}, {28, 8}, {28, 9}, {17, 2}, {17, 1}, {16, 1}, {15, 1}, {14, 1}, {13, 1}, {12, 1}, {11, 1}, {18, 1}, {19, 1}, {20, 1}, {21, 1}, {22, 1}, {23, 1}},
		},
		{
			traits: []string{"Blue Sunglasses", "Red Sunglasses", "Green Sunglasses"},
			maskLocationsToUnset: [][]int{{29, 13}, {29, 14}},
		},
		{
			traits: []string{"Headphones"},
			maskLocationsToUnset: [][]int{{9, 6}, {10, 6}, {10, 5}, {10, 4}, {11, 5}, {11, 4}, {11, 3}, {12, 4}, {12, 3}, {13, 4}, {13, 3}, {13, 2}, {14, 4}, {14, 3}, {14, 2}, {15, 4}, {15, 3}, {15, 2}, {16, 3}, {16, 2}, {17, 3}, {17, 2}, {18, 3}, {18, 2}, {19, 3}, {19, 2}, {20, 3}, {20, 2}, {21, 3}},
		},
		{
			traits: []string{"Flat Top"},
			maskLocationsToUnset: [][]int{{4, 12}, {5, 10}, {6, 8}, {7, 7}, {7, 6}, {8, 7}, {8, 6}, {8, 5}, {9, 6}, {9, 5}, {9, 4}, {10, 6}, {10, 5}, {10, 4}, {10, 3}, {11, 5}, {11, 4}, {11, 3}, {11, 2}, {12, 4}, {12, 3}, {12, 2}, {13, 4}, {13, 3}, {13, 2}, {13, 1}, {14, 4}, {14, 3}, {14, 2}, {14, 1}, {15, 4}, {15, 3}, {15, 2}, {15, 1}, {16, 3}, {16, 2}, {16, 1}, {17, 3}, {17, 2}, {17, 1}, {18, 3}, {18, 2}, {18, 1}, {19, 3}, {19, 2}, {19, 1}, {20, 3}, {20, 2}, {20, 1}, {21, 3}, {21, 2}, {21, 1}, {22, 3}, {22, 2}, {22, 1}, {23, 4}, {23, 3}, {23, 2}, {23, 1}, {24, 4}, {24, 3}, {24, 2}, {24, 1}, {25, 5}, {25, 4}, {25, 3}, {25, 2}, {25, 1}, {26, 7}, {26, 6}, {26, 5}, {26, 4}, {26, 3}, {26, 2}, {26, 1}, {27, 8}, {27, 7}, {27, 6}, {27, 5}, {27, 4}, {28, 9}, {28, 8}, {28, 7}},
		},
		{
			traits: []string{"Purple Bandana", "Blue Bandana", "Yellow Bandana", "Green Bandana"},
			maskLocationsToUnset: [][]int{{2, 11}, {3, 11}, {3, 10}, {4, 10}, {5, 10}, {5, 9}, {5, 8}, {4, 8}, {3, 8}, {2, 8}, {2, 7}, {3, 7}, {6, 8}, {7, 7}, {8, 7}, {9, 6}, {10, 6}, {10, 5}, {11, 5}, {11, 4}, {12, 4}, {13, 4}, {13, 3}, {14, 4}, {14, 3}, {15, 4}, {15, 3}, {16, 3}, {17, 3}, {18, 3}, {19, 3}, {20, 3}, {21, 3}, {22, 3}, {23, 4}, {24, 4}, {25, 5}, {26, 6}, {26, 7}, {27, 8}},
		},
		{
			traits: []string{"Orange Beanie", "Blue Beanie", "Green Beanie", "Grey Beanie"},
			maskLocationsToUnset: [][]int{{5, 10}, {5, 9}, {5, 8}, {6, 8}, {6, 7}, {7, 7}, {7, 6}, {8, 7}, {8, 6}, {8, 5}, {9, 6}, {9, 5}, {9, 4}, {10, 6}, {10, 5}, {10, 4}, {11, 5}, {11, 4}, {11, 3}, {12, 4}, {12, 3}, {13, 4}, {13, 3}, {14, 4}, {14, 3}, {15, 4}, {15, 3}, {16, 3}, {17, 3}, {18, 3}, {19, 3}, {20, 3}, {21, 3}, {22, 3}, {23, 4}, {24, 4}, {25, 5}, {26, 7}, {26, 6}, {27, 8}, {27, 7}, {28, 9}, {28, 8}},
		},
		{
			traits: []string{"CK Cap"},
			maskLocationsToUnset: [][]int{{5, 10}, {5, 9}, {6, 8}, {7, 7}, {8, 7}, {8, 6}, {9, 6}, {9, 5}, {10, 6}, {10, 5}, {10, 4}, {11, 5}, {11, 4}, {12, 4}, {13, 4}, {14, 4}, {14, 3}, {15, 4}, {15, 3}, {16, 3}, {17, 3}, {18, 3}, {19, 3}, {19, 2}, {20, 3}, {20, 2}, {20, 1}, {21, 3}, {21, 2}, {21, 1}, {22, 3}, {22, 2}, {22, 1}, {23, 4}, {23, 3}, {23, 2}, {23, 1}, {24, 4}, {24, 3}, {24, 2}, {24, 1}, {25, 5}, {25, 4}, {25, 3}, {25, 2}, {25, 1}, {26, 7}, {26, 6}, {26, 5}, {26, 4}, {26, 3}, {26, 2}, {26, 1}, {27, 8}, {27, 7}, {27, 6}, {27, 5}, {27, 4}, {27, 3}, {27, 2}, {28, 9}, {28, 8}, {28, 7}, {28, 6}, {28, 5}, {28, 4}, {29, 10}, {29, 9}, {29, 8}, {29, 7}},
		},
		{
			traits: []string{"Mohawk"},
			maskLocationsToUnset: [][]int{{15, 4}, {17, 3}, {18, 3}, {19, 3}, {20, 3}, {20, 2}, {21, 3}, {21, 2}, {22, 3}, {22, 2}, {22, 1}, {23, 4}, {23, 3}, {23, 2}, {23, 1}},
		},
		{
			traits: []string{"E-Cigarette"},
			maskLocationsToUnset: [][]int{{29, 25}, {29, 26}, {30, 24}, {30, 25}, {30, 26}, {31, 25}},
		},
		{
			traits: []string{"Green Bowtie", "Red Bowtie"},
			maskLocationsToUnset: [][]int{{29, 28}, {30, 29}},
		},
		{
			traits: []string{"Afro"},
			maskLocationsToUnset: [][]int{{0, 20}, {0, 19}, {0, 18}, {0, 17}, {0, 16}, {0, 15}, {0, 14}, {0, 13}, {0, 12}, {0, 11}, {0, 10}, {1, 16}, {1, 15}, {1, 14}, {1, 13}, {1, 12}, {1, 11}, {1, 10}, {1, 9}, {1, 8}, {2, 14}, {2, 13}, {2, 12}, {2, 11}, {2, 10}, {2, 9}, {2, 8}, {2, 7}, {2, 6}, {3, 13}, {3, 12}, {3, 11}, {3, 10}, {3, 9}, {3, 8}, {3, 7}, {3, 6}, {3, 5}, {3, 4}, {4, 12}, {4, 11}, {4, 10}, {4, 9}, {4, 8}, {4, 7}, {4, 6}, {4, 5}, {4, 4}, {4, 3}, {5, 10}, {5, 9}, {5, 8}, {5, 7}, {5, 6}, {5, 5}, {5, 4}, {5, 3}, {6, 8}, {6, 7}, {6, 6}, {6, 5}, {6, 4}, {6, 3}, {6, 2}, {7, 7}, {7, 6}, {7, 5}, {7, 4}, {7, 3}, {7, 2}, {8, 7}, {8, 6}, {8, 5}, {8, 4}, {8, 3}, {8, 2}, {9, 6}, {9, 5}, {9, 4}, {9, 3}, {9, 2}, {9, 1}, {10, 6}, {10, 5}, {10, 4}, {10, 3}, {10, 2}, {10, 1}, {10, 0}, {11, 5}, {11, 4}, {11, 3}, {11, 2}, {11, 1}, {11, 0}, {12, 4}, {12, 3}, {12, 2}, {12, 1}, {12, 0}, {13, 4}, {13, 3}, {13, 2}, {13, 1}, {13, 0}, {14, 4}, {14, 3}, {14, 2}, {14, 1}, {14, 0}, {15, 4}, {15, 3}, {15, 2}, {15, 1}, {15, 0}, {16, 3}, {16, 2}, {16, 1}, {16, 0}, {17, 3}, {17, 2}, {17, 1}, {17, 0}, {18, 3}, {18, 2}, {18, 1}, {18, 0}, {19, 3}, {19, 2}, {19, 1}, {19, 0}, {20, 3}, {20, 2}, {20, 1}, {20, 0}, {21, 3}, {21, 2}, {21, 1}, {21, 0}, {22, 4}, {22, 3}, {22, 2}, {22, 1}, {22, 0}, {23, 5}, {23, 4}, {23, 3}, {23, 2}, {23, 1}, {23, 0}, {24, 5}, {24, 4}, {24, 3}, {24, 2}, {24, 1}, {24, 0}, {25, 6}, {25, 5}, {25, 4}, {25, 3}, {25, 2}, {25, 1}, {25, 0}, {26, 7}, {26, 6}, {26, 5}, {26, 4}, {26, 3}, {26, 2}, {26, 1}, {27, 8}, {27, 7}, {27, 6}, {27, 5}, {27, 4}, {27, 3}, {27, 2}, {27, 1}, {28, 12}, {28, 9}, {28, 8}, {28, 7}, {28, 6}, {28, 5}, {28, 4}, {28, 3}, {28, 2}, {29, 14}, {29, 13}, {29, 12}, {29, 11}, {29, 10}, {29, 9}, {29, 8}, {29, 7}, {29, 6}, {29, 5}, {29, 4}, {29, 3}, {29, 2}, {30, 13}, {30, 12}, {30, 11}, {30, 10}, {30, 9}, {30, 8}, {30, 7}, {30, 6}, {30, 5}, {30, 4}, {30, 3}, {31, 9}, {31, 8}, {31, 7}, {31, 6}, {31, 5}},
		},
		{
			traits: []string{"Head Mirror"},
			maskLocationsToUnset: [][]int{{6, 8}, {7, 7}, {8, 7}, {19, 3}, {20, 3}, {20, 2}, {21, 3}, {21, 2}, {22, 3}, {22, 2}, {23, 4}, {23, 3}, {23, 2}, {24, 4}, {24, 3}, {25, 5}, {25, 4}, {26, 7}, {26, 6}, {26, 5}, {27, 8}, {28, 9}},
		},
		{
			traits: []string{"Hard Hat"},
			maskLocationsToUnset: [][]int{{5, 10}, {5, 9}, {6, 8}, {7, 7}, {8, 7}, {8, 6}, {8, 5}, {9, 6}, {9, 5}, {9, 4}, {10, 6}, {10, 5}, {10, 4}, {10, 3}, {11, 5}, {11, 4}, {11, 3}, {12, 4}, {12, 3}, {12, 2}, {13, 4}, {13, 3}, {13, 2}, {14, 4}, {14, 3}, {14, 2}, {15, 4}, {15, 3}, {15, 2}, {15, 1}, {16, 3}, {16, 2}, {16, 1}, {17, 3}, {17, 2}, {17, 1}, {18, 3}, {18, 2}, {18, 1}, {19, 3}, {19, 2}, {19, 1}, {20, 3}, {20, 2}, {20, 1}, {21, 3}, {21, 2}, {22, 3}, {22, 2}, {23, 4}, {23, 3}, {23, 2}, {24, 4}, {24, 3}, {25, 5}, {25, 4}, {26, 7}, {26, 6}, {26, 5}, {27, 8}, {27, 7}, {27, 6}, {28, 9}, {28, 8}, {28, 7}, {29, 10}, {29, 9}, {29, 8}, {29, 7}, {30, 10}, {30, 9}, {30, 8}, {30, 7}, {31, 9}, {31, 8}},
		},
		{
			traits: []string{"Straw Hat"},
			maskLocationsToUnset: [][]int{{1, 10}, {2, 10}, {3, 10}, {4, 10}, {5, 10}, {5, 9}, {5, 8}, {6, 8}, {6, 7}, {6, 6}, {7, 7}, {7, 6}, {7, 5}, {8, 7}, {8, 6}, {8, 5}, {8, 4}, {9, 6}, {9, 5}, {9, 4}, {9, 3}, {10, 6}, {10, 5}, {10, 4}, {10, 3}, {11, 5}, {11, 4}, {11, 3}, {11, 2}, {12, 4}, {12, 3}, {12, 2}, {13, 4}, {13, 3}, {13, 2}, {14, 4}, {14, 3}, {14, 2}, {14, 1}, {15, 4}, {15, 3}, {15, 2}, {15, 1}, {16, 3}, {16, 2}, {16, 1}, {17, 3}, {17, 2}, {17, 1}, {18, 3}, {18, 2}, {18, 1}, {19, 3}, {19, 2}, {19, 1}, {20, 3}, {20, 2}, {21, 3}, {21, 2}, {22, 3}, {22, 2}, {23, 4}, {23, 3}, {24, 4}, {24, 3}, {25, 5}, {25, 4}, {26, 7}, {26, 6}, {26, 5}, {27, 8}, {27, 7}, {27, 6}, {28, 8}, {28, 9}, {29, 10}, {30, 10}, {31, 10}, {32, 10}},
		},
		{
			traits: []string{"Brain"},
			maskLocationsToUnset: [][]int{{14, 4}, {15, 4}},
		},
		{
			traits: []string{"VR Goggles"},
			maskLocationsToUnset: [][]int{{29, 14}, {29, 15}, {29, 16}, {29, 17}},
		},
	}

	// Apply the mask
	for _, kongzTrait := range kongzTraits {
		for _, traitsSubstractiveMask := range traitsSubstractiveMasks {
			for _, traitOfSubstractiveMask := range traitsSubstractiveMask.traits {
				if kongzTrait == traitOfSubstractiveMask {
					// Ok our kongz needs to get this substractive mask
					for _, maskLocationToUnset := range traitsSubstractiveMask.maskLocationsToUnset {
						mask[maskLocationToUnset[1]][maskLocationToUnset[0]] = 0
					}
				}
			}
		}
	}

	return
}


type Attribute struct {
	TraitType string `json:"trait_type"`
	Value string `json:"value"`
}

type Metadata struct {
	Image string `json:"image"`
	ExternalUrl string `json:"external_url"`
	Name string `json:"name"`
	Attributes []Attribute `json:"attributes"`
}

func getKongzTraitByType(traitType string, id uint) (traits []string) {
	// Read metadata
	jsonFile, err := os.Open(fmt.Sprintf("assets/metadata/%d", id))
	if err != nil {
	    fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var metadata Metadata
	json.Unmarshal(byteValue, &metadata)

	for _, attribute := range metadata.Attributes {
		if attribute.TraitType == traitType {
			traits = append(traits, attribute.Value)
		}
	}

	return
}