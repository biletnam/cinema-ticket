package db

import(
  //"fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson")

type Time struct{
  Hour string `json: hour`
  Seats [][]bool `json: seats`
}

type Movie struct{
  Movie_name string `json: movie_name`
  Movie_img string `json: movie_img`
  Screen string `json: screen`
  Synopsis string `json: synopsis`
  Times []Time `json: times`
}

func GetMovies() ([]Movie,error){
  session, err := mgo.Dial("localhost:27017")

  if err != nil{
    return nil,err
  }
  defer session.Close()

  c := session.DB("cinema").C("movies")

  result := []Movie{}

  err = c.Find(bson.M{}).All(&result)

  if err != nil{
    return nil, err
  }

  return result, nil
}