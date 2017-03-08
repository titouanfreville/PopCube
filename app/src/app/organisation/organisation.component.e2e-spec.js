describe('Server', function () {

  beforeEach(function() {
      element(by.css('my-app header nav a:last-child')).click();
  });

  it('should have <my-server>', function () {
    var home = element(by.css('my-app my-server'));
    expect(home.isPresent()).toEqual(true);
  });

});
