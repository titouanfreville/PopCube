describe('Login', function () {

  beforeEach(function() {
      element(by.css('my-app header nav a:last-child')).click();
  });

  it('should have <my-login>', function () {
    var home = element(by.css('my-app my-login'));
    expect(home.isPresent()).toEqual(true);
  });

});
