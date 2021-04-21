Feature: Suministrar Jarabe a una mascota

  Scenario: suministrar Jarabe de forma extosa para el rango uno 
    Given el listado de mascotas tiene 3 mascotas
    When existen mascotas en el listado de mascotas 
    And y se envia el id de la mascota: "12NG"
    Then el resultado es la cantidad de CC a suministrar = 3

    Scenario: suministrar Jarabe de forma extosa para el rango 2 
    Given el listado de mascotas tiene 3 mascotas
    When existen mascotas en el listado de mascotas 
    And y se envia el id de la mascota: "WG15"
    Then el resultado es la cantidad de CC a suministrar = 12.5

    Scenario: suministrar Jarabe de forma extosa para el rango 3 
    Given el listado de mascotas tiene 3 mascotas
    When existen mascotas en el listado de mascotas 
    And y se envia el id de la mascota: "89PE"
    Then el resultado es la cantidad de CC a suministrar = 3

  Scenario: Error por que el Id de la mascota no existe
    Given el listado de mascotas tiene 3 mascotas
    When existen mascotas en el listado de mascotas 
    And y se envia el id de la mascota: "WG22"
    Then el resultado es la cantidad de CC a suministrar = 0
    And genera error "No se Puede suministrar Jarabe: la Mascota no existe"

  Scenario: Error por que la edad de la mascota es menor a la minima para el Jarabe
    Given el listado de mascotas tiene 3 mascotas
    When existen mascotas en el listado de mascotas 
    And y se envia el id de la mascota: "WG15" con edad = 1
    Then el resultado es la cantidad de CC a suministrar = 0
    And genera error "No se Puede suministrar Jarabe: La mascota no tiene la edad solicitada"
    
  
  