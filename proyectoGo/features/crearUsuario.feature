Feature: crear usuario

  Scenario: crear un usuario de forma exitosa
    Given el listado de mascotas tiene 3 mascotas
    And  el listado de vacunas tiene 2 vacunas
    When existe la mascota y la vacuna en el listado de mascotas y vacunas
    And y se envia la vacuna id 1 y el id de la mascota: "WG15"
    Then el resultado es la cantidad de CC a aplicar = 11.5

  Scenario: Error por que el Id de la mascota no existe
    Given el listado de mascotas tiene 3 mascotas
    And  el listado de vacunas tiene 2 vacunas
    When existe la mascota y la vacuna en el listado de mascotas y vacunas
    And y se envia la vacuna id 1 y el id de la mascota: "WG22"
    Then el resultado es la cantidad de CC a aplicar = 0
    And genera error "No se Puede aplicar la Vacuna: la Mascota no existe"

  Scenario: Error por que el Id de la vacuna no existe
    Given el listado de mascotas tiene 3 mascotas
    And  el listado de vacunas tiene 2 vacunas
    When existe la mascota y la vacuna en el listado de mascotas y vacunas
    And y se envia la vacuna id 3 y el id de la mascota: "WG15"
    Then el resultado es la cantidad de CC a aplicar = 0
    And genera error "No se Puede aplicar la Vacuna: la Vacuna no existe"

  Scenario: Error por que la vacuna no es para el tipo de mascota
    Given el listado de mascotas tiene 3 mascotas
    And  el listado de vacunas tiene 2 vacunas
    When existe la mascota y la vacuna en el listado de mascotas y vacunas
    And y se envia la vacuna id 1 y el id de la mascota: "12NG"
    Then el resultado es la cantidad de CC a aplicar = 0
    And genera error "No se Puede aplicar la Vacuna: No es para este tipo de mascota"

  Scenario: Error por que la edad de la mascota es menor a la minima de la vacuna
    Given el listado de mascotas tiene 3 mascotas
    And  el listado de vacunas tiene 2 vacunas
    When existe la mascota y la vacuna en el listado de mascotas y vacunas
    And y se envia la vacuna id 1 y el id de la mascota: "WG15" con edad = 1
    Then el resultado es la cantidad de CC a aplicar = 0
    And genera error "No se Puede aplicar la Vacuna: La mascota no tiene la edad solicitada"
    
  
  