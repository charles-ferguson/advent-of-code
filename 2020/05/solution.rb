#!/usr/bin/env ruby

DATA =
  begin
    file = File.read(File.join(__dir__, 'data'))
    file.lines.map(&:chomp)
  end

class Seat
  ROW_IDENTIFIER_LENGTH = 6
  COLUMN_IDENTIFIER_LENGTH = 2

  attr_reader :column_partition, :row_partition, :partition
  def initialize(binary_partition)
    @partition = binary_partition
    @row_partition = binary_partition[0..ROW_IDENTIFIER_LENGTH].reverse
    @column_partition = binary_partition.reverse[0..COLUMN_IDENTIFIER_LENGTH]
  end

  def seat_id
    row * 8 + column
  end

  def output
    "#{partition}:  row #{row},  column #{column},  seat ID #{seat_id}."
  end

  private

  def row
    @row ||=
      begin
        row_partition.chars.each_with_index.inject(0) do |row_count, (char, index)|
          char == 'B' ? row_count + 2**index : row_count
        end
      end
  end

  def column
    @column ||=
      begin
        column_partition.chars.each_with_index.inject(0) do |column_count, (char, index)|
          char == 'R' ? column_count + 2**index : column_count
        end
      end
  end
end

class SeatFinder
  MAX_ROW = (0..Seat::ROW_IDENTIFIER_LENGTH).inject(0) { |sum, index| sum + 2 ** index }
  MAX_COLUMN = (0..Seat::COLUMN_IDENTIFIER_LENGTH).inject(0) { |sum, index| sum + 2 ** index }

  attr_accessor :possible_seats
  def initialize
    @possible_seats = (0..(MAX_ROW * (MAX_COLUMN + 1) + MAX_COLUMN)).to_a
  end

  def filter_seats(seat_ids)
    self.possible_seats -= seat_ids
    self
  end

  def filter_row(row_number)
    self.possible_seats -= ((row_number * (MAX_COLUMN + 1))..(row_number * (MAX_COLUMN + 1) + MAX_COLUMN)).to_a
    self
  end

  def filter_seats_with_neighbor
    self.possible_seats=
      possible_seats.partition { |seat| possible_seats.include?(seat -1) || possible_seats.include?(seat + 1) }.last
    self
  end
end

seats = DATA.map { |partition| Seat.new(partition) }
puts "Part 1: " + seats.max { |a, b| a.seat_id <=> b.seat_id }.output

seat_finder = SeatFinder.new
                        .filter_row(0)
                        .filter_row(SeatFinder::MAX_ROW)
                        .filter_seats(seats.map(&:seat_id))
                        .filter_seats_with_neighbor

puts "Part 2: remaining seat #{seat_finder.possible_seats.inspect}"
