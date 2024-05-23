import datetime

class Persona:
    def __init__(self, nombre:str, apellido:str, id:int, edad:int):
        self.nombre = nombre
        self.apellido = apellido
        self.id = id
        self.edad = edad
        pass

    def __str__(self):
        return f"Nombre: {self.nombre} {self.apellido}, ID: {self.id}, Edad: {self.edad}"

    def editNombre(self, n):
        self.nombre = n

    def editApellido(self, a):
        self.apellido = a

    def editID(self, i):
        self.id = i

    def editEdad(self, e):
        self.edad = e
