describe('Greetings', () => {
  it('should say hello after entering our name', () => {
    cy.visit('http://localhost:34115')
    
    cy.get('#input #name').type('AlbinoDrought')
    cy.get('#input button').click()

    cy.get('#result').should('contain', 'Hello AlbinoDrought, It\'s show time!')
  })
})