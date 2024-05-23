import datetime
from Evento import *
from Agenda import *
from Persona import *

class Contacto:
    def __init__(self, persona: Persona, mail: str, telefono: int):
        self.persona = persona
        self.mail = mail
        self.telefono = telefono
        pass

    def __str__(self):
        return f"Nombre: {self.persona.nombre} {self.persona.apellido}, mail: {self.mail}, teléfono: {self.telefono}"

    def printContacto (self):
        print(f"""
Nombre: {self.persona.nombre} {self.persona.apellido}
Correo: {self.mail}
Teléfono: {self.telefono}""")