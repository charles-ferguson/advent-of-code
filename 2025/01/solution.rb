#!/usr/bin/env ruby

# frozen_string_literal: true

require 'Logger'
LOGGER = Logger.new($stdout)

class Move
  attr_reader :direction, :steps

  def initialize(direction, steps)
    @direction = direction
    @steps = steps.to_i
  end
end

class Dial
  attr_reader :position, :max_position, :pointed_at_zero_times, :clicked_passed_zero_times

  def initialize(position = 50, max_position = 99)
    @max_position = max_position
    @position = 50
    @pointed_at_zero_times = position.zero? ? 1 : 0
    @clicked_passed_zero_times = position.zero? ? 1 : 0
  end

  def turn(move)
    case move.direction
    when 'L'
      @position -= move.steps
      zeros_passed, pos = @position.divmod(@max_position + 1)
      @clicked_passed_zero_times += zeros_passed.abs
      @position = pos
    when 'R'
      @position += move.steps
      zeros_passed, pos = @position.divmod(@max_position + 1)
      @clicked_passed_zero_times += zeros_passed
      @position = pos
    end

    @pointed_at_zero_times += 1 if @position.zero?
  end
end

def parse_input(file_path = 'puzzle_input.txt')
  lines = File.readlines(file_path, chomp: true)
  moves = []
  lines.each do |line|
    if line =~ /^(?<direction>L|R)(?<steps>\d+)$/
      moves << Move.new(Regexp.last_match[:direction], Regexp.last_match[:steps])
      LOGGER.debug "Parsed move: Direction=#{moves.last.direction}, Steps=#{moves.last.steps}"
    else
      LOGGER.warn "Invalid line format: #{line}"
    end
  end
  moves
end

dial = Dial.new
moves = parse_input
moves.each do |move|
  LOGGER.info "Executing move: Direction=#{move.direction}, Steps=#{move.steps}"
  dial.turn(move)
  LOGGER.info "Dial position: #{dial.position}, Times at zero: #{dial.pointed_at_zero_times}, Clicks passed zero: #{dial.clicked_passed_zero_times}"
end
