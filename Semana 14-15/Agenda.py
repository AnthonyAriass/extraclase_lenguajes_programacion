import datetime
from Contacto import *
from Evento import *
from Persona import *

class Agenda:
    def __init__(self, owner: Persona, contactos: list, eventos: list):
        self.contactos = contactos
        self.eventos = eventos
        self.owner = owner
        pass

    def printContactos(self):
        print("Contactos de", self.owner.nombre,":")
        for i in self.contactos:
            i.printContacto()

    def addContactos(self, contacto):
        self.contactos.append(contacto)

    def delContactos(self, nombre, apell):
        for i in self.contactos:
            if i.persona.nombre == nombre and i.persona.apellido == apell:
                self.contactos.remove(i)

    def modContacto(self, nombre, apell, cambios:list):
        for i in self.contactos:
            if i.persona.nombre == nombre and i.persona.apellido == apell:
                i.mail = cambios[0]
                i.telefono = cambios [1]

    def printEventos(self):
        if self.eventos != []:
            for i in self.eventos:
                print(i)
        else:
            print("No hay eventos para mostrar")
    
    def addEvento(self, ev):
        self.eventos.append(ev)
    
    def delEvento(self, nombre:str):
        for i in self.eventos:
            if i.nombre == nombre:
                self.eventos.remove(i)
    
    def modEvento (self, nombre, cambios:list):
        for i in self.eventos:
            if i.nombre == nombre:
                if cambios[0] != "":
                    i.nombre = cambios[0]
                if cambios[1] != "":
                    i.fecha = cambios[1]
                if cambios[2] != []:
                    i.invitados = cambios[2]

    def addInvitado(self, nombre:str, invitado):
        for i in self.eventos:
            if i.nombre == nombre:
                i.invitados.append(invitado)