Feature: Eliminar una mascota del listado de Mascotas

  Scenario: Eliminar correctamente una mascota enviando el Id de la Mascota
    Given el listado de mascotas tiene 3 mascotas
    When se va a eliminar del listado la mascota con id: "12NG"
    Then el resultado es un listado con 2 mascotas

  Scenario: Error por que el Id de la mascota no existe
    Given el listado de mascotas tiene 3 mascota
    When  se va a eliminar del listado la mascota con id: "WG22" 
    Then el resultado es un listado con 3 mascotas
    And genera error "Error la mascota no existe"
    
  
  