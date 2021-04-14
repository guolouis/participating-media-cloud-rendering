package animations

import (
    "fmt"
    "math"
    "strconv"
    "volumetric-cloud/camera"
    "volumetric-cloud/ray"
    "volumetric-cloud/vector3"
    "volumetric-cloud/scene"
    "volumetric-cloud/img"

)

func AnimRotation(center vector3.Vector3, radius float64,
    imgSizeX, imgSizeY, picNumber int, s scene.Scene){

    s.Camera.Origin.X = 0.0
    s.Camera.Origin.Y = 0.0
    s.Camera.Origin.Z = 0.0
    s.Camera.RotationX, s.Camera.RotationY, s.Camera.RotationZ =
        camera.InitRota(0.0, 0.0, 0.0)
    for i := 0; i < picNumber + 1; i+=1 {
        teta := (float64(i) * 2 * math.Pi) / float64(picNumber)
        s.Camera.Origin.X = -radius * math.Sin(teta) + center.X
        s.Camera.Origin.Y = center.Y
        s.Camera.Origin.Z = radius * math.Cos(teta) + center.Z


        s.Camera.RotationX, s.Camera.RotationY, s.Camera.RotationZ =
                    camera.InitRota(0.0, - teta, 0.0)

        image := s.Render(imgSizeX, imgSizeY)
        image.SavePNG("videos/video_img" + strconv.Itoa(i) + ".png")
        s.Pixels = s.Pixels[:0]
        fmt.Println("---- img" + strconv.Itoa(i) + "---- done")
    }
}

func AnimTranslate(ray ray.Ray, picNumber, imgSizeX, imgSizeY int,
    step float64, s scene.Scene, cam camera.Camera){

    for i:=0; i < picNumber; i++ {
        newPos := ray.RayAt(step * float64(i))
        s.Camera.Origin.X = newPos.X
        s.Camera.Origin.Y = newPos.Y
        s.Camera.Origin.Z = newPos.Z

        image := s.Render(imgSizeX, imgSizeY)
        image.SavePNG("videos/video_img" + strconv.Itoa(i) + ".png")
        s.Pixels = s.Pixels[:0]
        fmt.Println("---- img " + strconv.Itoa(i) + "---- done")
    }
}

func LookCenter(point vector3.Vector3, s scene.Scene, imgSizeX, imgSizeY int) img.Img {

    teta := math.Atan((point.X - s.Camera.Origin.X) / (point.Z - s.Camera.Origin.Z))
    phi := math.Atan((point.Y - s.Camera.Origin.Y) / (point.Z - s.Camera.Origin.Z))

    //teta += 2.0 * math.Pi / float64(picNumber)
    s.Camera.RotationX, s.Camera.RotationY, s.Camera.RotationZ =
        camera.InitRota(phi, teta, 0.0)

    return s.Render(imgSizeX, imgSizeY)
}





