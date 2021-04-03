package scene

import (
    "fmt"
    "os"
    "volumetric-cloud/img"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/camera"
    "volumetric-cloud/sphere"
    "volumetric-cloud/light"
    "volumetric-cloud/background"
    "volumetric-cloud/vector3"

    "math/rand"
    "sync"
)

type Scene struct {
    // TODO change to many VoxelGrid
    VoxelGrids []voxel_grid.VoxelGrid
    Sphere sphere.Sphere
    Camera camera.Camera
    Light light.Light
}

func InitScene(voxelGrids []voxel_grid.VoxelGrid,
               sphere sphere.Sphere,
               camera camera.Camera,
               light light.Light) Scene {
    // compute light transmittance in the voxel grid
    for idx, _ := range voxelGrids {
        voxelGrids[idx].ComputeInsideLightTransparency(light)
    }

    return Scene{
        VoxelGrids: voxelGrids,
        Sphere: sphere,
        Camera: camera,
        Light: light,
    }
}

func (s Scene) Render(imgSizeY, imgSizeX, nbRaysPerPixel int) img.Img {
    image := img.InitImg(imgSizeY, imgSizeX)

    // create the wait group
    wg := sync.WaitGroup{}
    wg.Add(imgSizeY)

    for i := 0; i < imgSizeY; i += 1 {
//        if i != 568 {
//            continue
//        }
        s.renderImageSizeY(image, i, imgSizeX, nbRaysPerPixel, nil)

        /*
        for j := 0; j < imgSizeX; j += 1 {
            //go s.renderPixel(image, i, j, nbRaysPerPixel, &wg)
            s.renderPixel(image, i, j, nbRaysPerPixel, nil)
        }
        */

    }
//    wg.Wait()
    return image
}



func (s Scene) renderImageSizeY(image img.Img, i, imgSizeX, nbRaysPerPixel int, wg *sync.WaitGroup) {
    for j := 0; j < imgSizeX; j += 1 {
//        if j != 482 {
//            continue
//        }
        color := vector3.InitVector3(0, 0, 0)
        for k := 0; k < nbRaysPerPixel; k += 1 {
            // create the ray
            r := s.Camera.CreateRay(float64(j) + rand.Float64(), float64(i) + rand.Float64())

            var accColor vector3.Vector3
            var accC vector3.Vector3

            var accTransparency float64 = 1.0
            var accT float64

            var hasHit bool
            var hasOneHit bool = false

            // Check intersect with Voxel Grids
            sum := 0
            for _, vGrid := range s.VoxelGrids {
                accC, accT, hasHit = vGrid.ComputePixelColor(r, s.Light.Color)
                if !hasHit {
                    continue
                }

                hasOneHit = true

                // accumulate transparency
                accTransparency *= accT

                // accumulate color
                accColor.AddVector3(accC)
                sum += 1
            }

            if sum == 2 {
                fmt.Printf("%v %v\n", i, j)
                os.Exit(1)
            }

            // get background impact
            backgroundColor := background.RenderGradient(r)

            // set pixel
            if hasOneHit {
                // compute pizel color
                backgroundColorImpact := vector3.MulVector3Scalar(backgroundColor, accTransparency)
                accColor.AddVector3(backgroundColorImpact)
                accColor.Clamp(0.0, 1.0)
                color.AddVector3(vector3.InitVector3(accColor.X, accColor.Y, accColor.Z))
            } else {
                // gradient case
                color.AddVector3(backgroundColor)
            }
        }

        // divide color vector by nbRaysPerPixel
        color.Div(float64(nbRaysPerPixel))

        // Set the pixel color
        image.SetPixel(i, j, byte(color.X * 255.0), byte(color.Y * 255.0), byte(color.Z * 255.0))


    }

    if wg != nil {
        wg.Done()
    }
}

/*
func (s Scene) renderPixel(image img.Img, i, j, nbRaysPerPixel int, wg *sync.WaitGroup) {
    color := vector3.InitVector3(0, 0, 0)
    for k := 0; k < nbRaysPerPixel; k += 1 {
        // create the ray
        r := s.Camera.CreateRay(float64(j) + rand.Float64(), float64(i) + rand.Float64())

        var accColor vector3.Vector3
        var accC vector3.Vector3

        var accTransparency float64 = 1.0
        var accT float64

        var hasHit bool
        var hasOneHit bool = false

        // Check intersect with Voxel Grids
        for _, vGrid := range s.VoxelGrids {
            accC, accT, hasHit = vGrid.ComputePixelColor(r, s.Light.Color)
            if !hasHit {
                continue
            } else {
                hasOneHit = true
            }

            // accumulate transparency
            accTransparency *= accT

            // accumulate color
            accColor.AddVector3(accC)
        }

        // get background impact
        backgroundColor := background.RenderGradient(r)

        // set pixel
        if hasOneHit {
            // compute pizel color
            backgroundColorImpact := vector3.MulVector3Scalar(backgroundColor, accTransparency)
            accColor.AddVector3(backgroundColorImpact)
            accColor.Clamp(0.0, 1.0)
            color.AddVector3(vector3.InitVector3(accColor.X, accColor.Y, accColor.Z))
        } else {
            // gradient case
            color.AddVector3(backgroundColor)
        }
    }

    // divide color vector by nbRaysPerPixel
    color.Div(float64(nbRaysPerPixel))

    // Set the pixel color
    image.SetPixel(i, j, byte(color.X * 255.0), byte(color.Y * 255.0), byte(color.Z * 255.0))

    if wg != nil {
        wg.Done()
    }
}

 */
