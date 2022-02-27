/// <reference types="cypress" />

describe('App', () => {
    beforeEach(() => {
        // Cypress starts out with a blank slate for each test
        // so we must tell it to visit our website with the `cy.visit()` command.
        // Since we want to visit the same URL at the start of all our tests,
        // we include it in our beforeEach function so that it runs before each test
        cy.visit('http://localhost:3000')
    })

    it('verify vendor login', () => {
        cy.get('#home-carousel').should('exist')
    })

    it('Verify vendor login', () => {
        cy.get('a[href="/vendorlogin"]').click()
        cy.get('input[id="emailId"]').should('exist')
        cy.get('input[id="passwordId"]').should('exist')
        cy.get('button[id="vendor-login-button"]').should('exist')
    })

    it('Verify user login', () => {
        cy.get('a[href="/login"]').click()
        cy.get('input[id="emailId"]').should('exist')
        cy.get('input[id="passwordId"]').should('exist')
        cy.get('button[id="user-login-button"]').should('exist')
    })

})