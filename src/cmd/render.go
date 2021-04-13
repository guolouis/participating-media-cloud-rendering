package cmd

import (
    "math"
    "fmt"

    "github.com/spf13/cobra"

    "volumetric-cloud/camera"
    "volumetric-cloud/light"
    "volumetric-cloud/noise"
    "volumetric-cloud/scene"
    "volumetric-cloud/vector3"
    "volumetric-cloud/voxel_grid"
    "volumetric-cloud/random_clouds"
    "volumetric-cloud/atmosphere"
    "volumetric-cloud/sphere"
//    "volumetric-cloud/animations"
)

var fullRenderCmd = &cobra.Command{
    Use: "fullrender",
    Short: "Generate clouds and render them",
    Run: func(cmd *cobra.Command, args []string) {
        imgSizeX := 1200
        imgSizeY := 1000

        // Camera
        aspectRatio := float64(imgSizeX) / float64(imgSizeY)
        fieldOfView := math.Pi / 2

        origin := vector3.InitVector3(0.0, 0.0, 5)
        camera := camera.InitCamera(
           aspectRatio,
           fieldOfView,
           imgSizeX,
           imgSizeY,
           origin,
           0.0,
           0.0,
           0.0,
        )

        shift := vector3.InitVector3(-20.0, 25.0, -90.0)
        oppositeCorner := vector3.InitVector3(20.0, 30.0, -50.0)
        var seed2 int64 = 2
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed2)
        worleyWeight := 0.2
        perlinWeight := 0.6
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise2, worleyNoise2, perlinWeight, worleyWeight, 0.3, 0.6, 1.5)

        // Voxel Grid 1
        // Image 'perlin-worley-4.png'
        /*shift := vector3.InitVector3(-20.0, 25.0, -100.0)
        oppositeCorner := vector3.InitVector3(20.0, 30.0, -50.0)
        var seed2 int64 = 21
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed2)
        worleyWeight := 0.4
        perlinWeight := 0.6
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift, oppositeCorner, 0.13, perlinNoise2, worleyNoise2, perlinWeight, worleyWeight, 0.6, 0.6, 1.5)*/

        // Voxel Grid 2
        // Image 'perlin-worley-2.png'
        /*shift2 := vector3.InitVector3(-50, 35.0, -60.0)
        oppositeCorner2 := vector3.InitVector3(-25.0, 40.0, -30.0)
        var seed2 int64 = 21
        worleyNoise2 := noise.InitWorleyNoise(0.4, 2.0, 0.5, 0.5, 3, seed2)
        perlinNoise2 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 3, seed2)
        worleyWeight := 0.5
        perlinWeight := 0.5
        voxelGrid2 := voxel_grid.InitVoxelGrid(0.5, shift2, oppositeCorner2, 0.13, perlinNoise2, worleyNoise2, perlinWeight, worleyWeight, 0.6, 0.6, 1.5)
        */

        // Voxel Grid 3
        // Image 'perlin-worley-1.png'
        /*shift3 := vector3.InitVector3(15.0, 30.0, -80.0)
        oppositeCorner3 := vector3.InitVector3(60.0, 38.0, -30.0)
        var seed3 int64 = 39
        worleyNoise3 := noise.InitWorleyNoise(0.2, 2.5, 0.5, 0.2, 2, seed3)
        perlinNoise3 := noise.InitPerlinNoise(0.2, 2.0, 1.0, 0.8, 4, seed3)
        worleyWeight := 0.4
        perlinWeight := 0.6
        voxelGrid3 := voxel_grid.InitVoxelGrid(0.5, shift3, oppositeCorner3, 0.13, perlinNoise3, worleyNoise3, perlinWeight, worleyWeight, 0.3, 0.1, 2.5)
        */

        // Voxel Grid 4
/*        shift4 := vector3.InitVector3(-40.0, 30.0, -90.0)
        oppositeCorner4 := vector3.InitVector3(-10.0, 40.0, -60.0)

        var seed4 int64 = 301
        worleyNoise3 := noise.InitWorleyNoise(0.2, 2.0, 0.5, 0.2, 3, seed4)
        perlinNoise3 := noise.InitPerlinNoise(0.2, 2.0, 0.5, 0.8, 4, seed4)
        worleyWeight := 0.1
        perlinWeight := 0.9
        voxelGrid3 := voxel_grid.InitVoxelGrid(0.5, shift4, oppositeCorner4, 0.13, perlinNoise3, worleyNoise3, perlinWeight, worleyWeight, 0.6, 0.6, 1.5)
*/
/*
        // Voxel Grid 5
        shift5 := vector3.InitVector3(-20.0, 48.0, -70.0)
        oppositeCorner5 := vector3.InitVector3(15.0, 55.0, -30.0)
        var seed5 int64 = 39
        perlinNoise5 := noise.InitWorleyNoise(0.2, 2.0, 1.0, 0.3, 3, seed5)
        voxelGrid5 := voxel_grid.InitVoxelGrid(0.5, shift5, oppositeCorner5, 0.13, perlinNoise5, 0.6, 0.6, 1.8)
*/
   //     voxelGrids := []voxel_grid.VoxelGrid{voxelGrid, voxelGrid2, voxelGrid3, voxelGrid4, voxelGrid5}
        voxelGrids := []voxel_grid.VoxelGrid{voxelGrid2}

        // IMPORTANT
        //
        // First condition:
        // (oppositeCorner.X - shift.X) / voxelSize  => MUST BE AN INTEGER (NO FLOAT)
        // (oppositeCorner.Y - shift.Y) / voxelSize  => MUST BE AN INTEGER (NO FLOAT)
        // (oppositeCorner.Z - shift.Z) / voxelSize  => MUST BE AN INTEGER (NO FLOAT)
        //
        // Second condition:
        // shift.X < oppositeCorner.X &&
        // shift.Y < oppositeCorner.Y &&
        // shift.Z < oppositeCorner.Z
        fmt.Println("VOXEL")

        // Lights
        light1 := light.InitLight(vector3.InitVector3(0.0, 200.0, 200.0), vector3.InitVector3(0.8, 0.8, 0.8))
        light2 := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.7, 0.7, 0.7))
        lights := []light.Light{light1, light2}

        // Atmosphere
        ground := sphere.InitSphere(vector3.InitVector3(0.0, -150.0, 0.0), 150.0)
        groundColor := vector3.InitVector3(91.0 / 255.0, 113 / 255.0, 182.0 / 255.0)
        albedo := 0.9
        atmosphere := atmosphere.Atmosphere{
            Ground: ground,
            GroundColor: groundColor,
            GroundAlbedo: albedo,
        }

        // Scene
        fmt.Println("SCENE")
        s := scene.InitScene(voxelGrids, camera, lights, atmosphere, 0.3)

        fmt.Println("RENDER")

        /*animations.AnimRotation(vector3.InitVector3(-25.0, 35.0, -75.0),
                     70.0,
                     imgSizeX,
                     imgSizeY,
                     20,
                     1,
                     s)*/
        // Render
        image := s.Render(imgSizeY, imgSizeX)

        //fmt.Println("SAVE")
        // Save
        image.SavePNG("tmp.png")

    },
}

var randomRenderCmd = &cobra.Command{
    Use: "randomrender",
    Short: "Generate random clouds and render",
    Run: func(cmd *cobra.Command, args []string) {
        imgSizeX := 1200
        imgSizeY := 1000

        // Camera
        aspectRatio := float64(imgSizeX) / float64(imgSizeY)
        fieldOfView := math.Pi / 2

        origin := vector3.InitVector3(0, 15, 5)
        camera := camera.InitCamera(
           aspectRatio,
           fieldOfView,
           imgSizeX,
           imgSizeY,
           origin,
           math.Pi / 8,
           0.0,
           0.0,
        )

        // Init random voxelGrids
        voxelGrids := random_clouds.GenerateRandomClouds(7, 5)

        fmt.Println("VOXEL")

        // Lights
        light1 := light.InitLight(vector3.InitVector3(0.0, 200.0, -200.0), vector3.InitVector3(0.7, 0.7, 0.7))
        //light2 := light.InitLight(vector3.InitVector3(0.0, 0.0, 0.0), vector3.InitVector3(0.7, 0.7, 0.7))
        lights := []light.Light{light1}

        // Scene
        fmt.Println("SCENE")
        s := scene.InitScene(voxelGrids, camera, lights, atmosphere.Atmosphere{}, 1.0)

        fmt.Println("RENDER")
        // Render
        image := s.Render(imgSizeY, imgSizeX)

        fmt.Println("SAVE")
        // Save
        image.SavePNG("tmp.png")
    },
}

var loadRenderCmd = &cobra.Command{
    Use: "loadrender",
    Short: "Load clouds and render them",
    Args: cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        // load data
        // build the scene
        // launch render
    },
}

func init() {
    cmd.AddCommand(fullRenderCmd)
    cmd.AddCommand(loadRenderCmd)
    cmd.AddCommand(randomRenderCmd)
}
