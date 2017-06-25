describe('Organisation', function () {

  beforeEach(function() {
      element(by.css('my-app header nav a:last-child')).click();
  });

  it('should have <my-organisation>', function () {
    var home = element(by.css('my-app my-organisation'));
    expect(home.isPresent()).toEqual(true);
  });

});
