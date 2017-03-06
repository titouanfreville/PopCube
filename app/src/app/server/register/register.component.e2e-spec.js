describe('Register', function () {

  beforeEach(function() {
      element(by.css('my-app header nav a:first-child')).click();
  });

  it('should have <my-register>', function () {
    var register = element(by.css('my-app my-register'));
    expect(register.isPresent()).toEqual(true);
    expect(register.getText()).toEqual("Home Works!");
  });

});
