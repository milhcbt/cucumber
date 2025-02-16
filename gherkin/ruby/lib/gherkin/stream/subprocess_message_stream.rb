require 'open3'
require 'gherkin/stream/protobuf_message_stream'

module Gherkin
  module Stream
    class SubprocessMessageStream
      def initialize(gherkin_executable, paths, print_source, print_ast, print_pickles)
        @gherkin_executable, @paths, @print_source, @print_ast, @print_pickles = gherkin_executable, paths, print_source, print_ast, print_pickles
      end

      def messages
        args = [@gherkin_executable]
        args.push('--no-source') unless @print_source
        args.push('--no-ast') unless @print_ast
        args.push('--no-pickles') unless @print_pickles
        args = args.concat(@paths)
        stdin, stdout, stderr, wait_thr = Open3.popen3(*args)
        if(stdout.eof?)
          error = stderr.read
          raise error
        end
        ProtobufMessageStream.new(stdout).messages
      end
    end
  end
end
