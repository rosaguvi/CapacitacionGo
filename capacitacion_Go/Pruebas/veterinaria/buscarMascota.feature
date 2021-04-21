Feature: Buscar una o varias mascotas de una lista

  Scenario: buscar sin criterio
    Given el listado de mascotas tiene 3 mascotas
    When se busca sin criterio 
    Then el resultado es un listado con 3 mascotas

  Scenario: buscar con el tipo de mascota perro 
    Given el listado de mascotas tiene 3 mascotas
    When se busca con la palabra perro
    Then el resultado es un listado con 1 mascotas
    
  Scenario: buscar con el nombre de mascota Pelusa 
    Given el listado de mascotas tiene 3 mascotas
    When se busca con la palabra Pelusa
    Then el resultado es un listado con 1 mascotas
    
  Scenario: buscar con el Id de mascota 89PE 
    Given el listado de mascotas tiene 3 mascotas
    When se busca con la palabra Pelusa
    Then el resultado es un listado con 1 mascotas

  Scenario: buscar cuando la palabra ingresada coincide con el nombre, el tipo de mascota el el ID
    Given el listado de mascotas tiene 3 mascotas
    When se busca con la palabra pe
    Then el resultado es un listado con 3 mascotas

  Scenario: buscar cuando la palabra ingresada no coincida con ningun criterio
    Given el listado de mascotas tiene 3 mascotas
    When se busca con la palabra Mascota
    Then el resultado es un listado con 0 mascotas
    And genera error "No se encontro ningun Registro"