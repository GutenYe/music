A.Router.map  ->
  @resource "albums", ->
  @resource "songs", ->
      
A.AlbumsIndexRoute = Ember.Route.extend
  model: (params) ->
    A.Album.find()

A.SongsIndexRoute = Ember.Route.extend
  model: (params) ->
    A.Song.find()
