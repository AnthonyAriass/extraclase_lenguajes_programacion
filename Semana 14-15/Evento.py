import datetime
from Contacto import *
from Agenda import *
from Persona import *


class Evento:
    def __init__(self, nombre: str, fecha: datetime, invitados: list):
        self.nombre = nombre
        self.fecha = fecha
        self.invitados = invitados
        pass

    def __str__(self):
        invitados_info = '\n'.join([str(invitado) for invitado in self.invitados])
        return f"\nNombre: {self.nombre}, \nfecha: {self.fecha}, \nInvitados: \n{invitados_info}\n"

    