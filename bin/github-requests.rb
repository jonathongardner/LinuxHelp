require 'socket'
require 'io/console'
# require 'base64'
require 'net/http'
require 'json'
class GithubRequest
  attr_reader :uri
  attr_accessor :token
  attr_writer :username

  def initialize(url, token: nil, username: nil)
    @uri = URI(url)
    @token = token
    @username = username
  end

  #------attr------
  def uri=(url)
    @uri = URI(url)
  end

  def host
    self.uri.host
  end

  def username
    return @username if @username

    print('Github username? ')
    @username = gets.chomp
  end
  #------attr------

  def post_keys(key:, title: nil)
    uri.path += 'user/keys'
    request = _post_authorized_request
    request.body = { key: key, title: title }.compact.to_json
    _response(request)
    puts "Successfuly added public key"
  end

  def get_keys
    uri.path += "users/#{self.username}/keys"
    request = _get_request
    _json_response(request)
  end

  private
    def _post_authorized_request
      request = Net::HTTP::Post.new(uri)

      request['Accept'] = 'application/vnd.github.v3+json'
      if token
        puts "Using github token for authorization"
        request['Authorization'] = "token #{token}"
      else
        pword = IO::console.getpass('Password? ')
        request.basic_auth self.username, pword
        # headers['Authorization'] = "Basic #{Base64.encode64("#{args.username}:#{pword}")}"
      end
      request
    end

    def _get_request
      request = Net::HTTP::Get.new(uri)
      request['Accept'] = 'application/vnd.github.v3+json'
      request
    end

    def _response(request)
      response = Net::HTTP.start(uri.host, uri.port, use_ssl: true) do |http|
        puts "Requesting #{uri}..."
        http.request(request)
      end
      raise "ERROR: #{response.body}" unless response.is_a?(Net::HTTPSuccess)
      response
    end

    def _json_response(request)
      JSON.parse(_response(request).body)
    end
end
