import datetime
from Contacto import *
from Evento import *
from Agenda import *
from Persona import *

def crearAgenda(owner):
    return Agenda(owner, [], [])

def crearPersona(n:str, a:str, i:int, e:int):
    return Persona(n, a, i ,e)

def crearContacto(p:Persona, m:str, t:int):
    return Contacto(p, m, t)

def crearPersonas():
    return [
    crearPersona("Anthony", "Arias", 208110917, 23),
    crearPersona("María", "González", 309118765, 28),
    crearPersona("Carlos", "Martínez", 407112233, 35),
    crearPersona("Laura", "Fernández", 508114456, 22),
    crearPersona("David", "Pérez", 609113478, 30),
    crearPersona("Sofía", "López", 710115679, 25),
    crearPersona("Andrés", "Ramírez", 811116890, 27),
    crearPersona("Daniela", "Torres", 912117901, 32),
    crearPersona("Miguel", "Hernández", 101311218, 29),
    crearPersona("Valeria", "Ruiz", 110211529, 24),
    crearPersona("Jorge", "Mendoza", 120312639, 26)
]

def main():

    #Creamos una lista de personas
    personas = crearPersonas()

    #Creamos una agenda cuyo dueño es la persona en el índice 0 de la lista personas (Anthony)
    agendaAnthony = crearAgenda(personas[0])


    #Primero vamos a hacer una lista de correos y otra de teléfonos que vamos a usar para crear los contactos
    listaCorreos = [
        "mgonzales@mail.com",
        "cmartinez@mail.com",
        "lfernandez@mail.com",
        "dperez@mail.com",
        "slopez@mail.com",
        "aramirez@mail.com",
        "dtorres@mail.com",
        "mhernandez@mail.com",
        "vruiz@mail.com",
        "jmendoza@mail.com"
    ]

    listaTelefonos = [
        88888888,
        82828282,
        81475236,
        87965412,
        87954123,
        88998899,
        87878787,
        85858585,
        84563217,
        87986554
    ]

    #Hacemos que el resto de persona de la lista se vuelvan contactos de Anthony y los agregamos a la agenda
    for i in range (1, len(personas)):
        #addContactos es la función para añadir contactos
        agendaAnthony.addContactos(crearContacto(personas[i], listaCorreos[i-1], listaTelefonos[i-1]))

    
    #Mostramos los contactos de Anthony
    #agendaAnthony.printContactos()

    #Vamos a eliminar a un contacto usando nombre y apellido
    agendaAnthony.delContactos("Carlos", "Martínez")
    
    #agendaAnthony.printContactos()

    #Para modificar un contacto hay que dar el nombre, apellido, y una lista con los nuevos cambios (mail y telefono).
    agendaAnthony.modContacto("María", "González", ["mgonzalesmod@mail.com", 60606060])

    agendaAnthony.printContactos()

    print("--------------------------------------------------------")

    #Ahora vamos a crear un evento, para esto se necesita un nombre, un time.date, y una lista de invitados. y lo agregamos a la lista de eventos de la agenda
    #acá hacen falta validaciones como ver si el índice de los contactos no se excede, pero para mantener este ejercicio simple no voy a ahondar en este tipo de validaciones.
    agendaAnthony.addEvento(Evento("Carne asada", datetime.datetime(2024, 5, 24), [agendaAnthony.contactos[0], agendaAnthony.contactos[3], agendaAnthony.contactos[6]]))

    agendaAnthony.printEventos()

    #Borramos el evento
    #agendaAnthony.delEvento("Carne asada")
    #agendaAnthony.printEventos()

    #Modificamos el evento. si los cambios van vacíos no se actualizan, por ejemplo, acá solo vamos a actualizar el nombre
    agendaAnthony.modEvento("Carne asada", ["Reunión familiar", "", []])

    agendaAnthony.printEventos()


    #tambien podemos añadir invitados de forma individual
    agendaAnthony.addInvitado("Reunión familiar", agendaAnthony.contactos[4])

    agendaAnthony.printEventos()


if __name__ == "__main__":
    main()

