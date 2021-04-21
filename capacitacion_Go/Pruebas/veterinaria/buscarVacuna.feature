Feature: Buscar una o varias Vacunas de una lista

  Scenario: buscar sin datos o buscar todas
    Given el listado de vacunas tiene 2 vacunas
    When se busca sin criterio 
    Then el resultado es un listado con 2 vacunas

  Scenario: buscar con el tipo de mascota perro 
    Given el listado de vacunas tiene 2 vacunas
    When se busca con la palabra perro
    Then el resultado es un listado con 1 vacuna
    
  Scenario: buscar con el nombre de vacuna "vacuna 2" 
    Given el listado de vacunas tiene 2 vacunas
    When se busca con las palabras "vacuna 2"
    Then el resultado es un listado con 1 vacuna
    
  Scenario: buscar con el Id de vacuna "1" 
    Given el listado de vacunas tiene 2 vacunas
    When se busca con el id "1"
    Then el resultado es un listado con 1 vacuna

  Scenario: buscar cuando la palabra ingresada coincide con el nombre, y el ID
    Given el listado de vacunas tiene 2 vacunas
    When se busca con la palabra "2"
    Then el resultado es un listado con 2 vacunas

  Scenario: buscar cuando la palabra ingresada no coincida con ningun criterio
    Given el listado de vacunas tiene 2 vacunas
    When se busca con la palabra Mascota
    Then el resultado es un listado con 0 vacunas
    And genera error "No se encontro ningun Registro"