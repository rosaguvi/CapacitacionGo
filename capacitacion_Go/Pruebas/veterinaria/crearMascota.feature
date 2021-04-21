Feature: Agregar una mascota al listado de mascotas

  Scenario: crear una mascota con datos correctos
    Given el listado de mascotas tiene 1 mascota
    And se va a crear una mascota con nombre: "Mariposa", categoria: "Perro", edad: 2 meses y peso: 5 libras
    When se envian los datos de una mascota para ser agregada al listado
    Then el resultado es un listado con 2 mascotas

  Scenario: agregar una mascota Repetida
    Given el listado de mascotas tiene 1 mascota
    When se envian los datos de una mascota para ser agregada al listado
    Then el resultado es un listado con 1 mascota
    And genera error "Mascota ya Existe en el listado"
    
  Scenario: Agregar una mascota sin nombre
    Given el listado de mascotas tiene 1 mascotas
    When se envian los datos de una mascota para ser agregada al listado, pero con nombre en blanco
    Then el resultado es un listado con 1 mascotas
    And genera error "Datos incompletos: Falta el Nombre"

  Scenario: Agregar una mascota sin categoria
    Given el listado de mascotas tiene 1 mascotas
    When se envian los datos de una mascota para ser agregada al listado, pero con categoria en blanco
    Then el resultado es un listado con 1 mascotas
    And genera error "Datos incompletos: Falta la categoria"

  Scenario: Agregar una mascota sin peso
    Given el listado de mascotas tiene 1 mascotas
    When se envian los datos de una mascota para ser agregada al listado, pero con peso en blanco
    Then el resultado es un listado con 1 mascotas
    And genera error "Datos incompletos: Falta el peso"

  Scenario: Agregar una mascota sin edad
    Given el listado de mascotas tiene 1 mascotas
    When se envian los datos de una mascota para ser agregada al listado, pero con edad en blanco
    Then el resultado es un listado con 2 mascotas
    And genera error "Posible Error: La edad se crea en cero"
    
  