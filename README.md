This is the code for the instagram REST API using Goland and MongoDB. 

The code has been divided into three files- main.go, user.go under controllers file and user.go under models file.

                                                 MAIN.GO

It describes sets up the server and describes the route of the instagram requets.

                                           USER.GO under CONTROLLERS file

It decrbes the functions which can handle the get and put requets from the users. It connects to a MongoDB database, and if a database is not present it makes its one.

                                           USER.GO under CONTROLLERS file

It describes the schema that is used in the database along with the parameters. It basically describes the model.


The API has been tested with Postman. The results are given below:


                                                                      GET Request


   ![Screenshot (6490)](https://user-images.githubusercontent.com/61285140/136655634-e7771d8e-ec91-4ba4-9b01-78e20f4750e4.png)
                      
                      
                                                                        POST Request
                                                                        
                                                                        
   ![Screenshot (6491)](https://user-images.githubusercontent.com/61285140/136655658-f0b9e4b3-6c86-4b61-8979-591ff9036b52.png)
