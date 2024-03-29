class Admin::CardsController < ApplicationController
  before_action :require_admin!

  def index
    @cards = Card.all
  end

  def new
    @card = Card.new format: 'markdown', sort_order: 0
    render 'form'
  end

  def create
    @card = Card.new(card_params)

    if @card.save
      redirect_to admin_cards_path
    else
      render 'form'
    end
  end

  def edit
    @card = Card.find(params[:id])
    render 'form'
  end

  def update
    @card = Card.find(params[:id])

    if @card.update(card_params)
      redirect_to admin_cards_path
    else
      render 'form'
    end
  end

  def destroy
    @card = Card.find(params[:id])
    @card.destroy

    redirect_to admin_cards_path
  end


  private
  def card_params
    params.require(:card).permit(:loc, :title, :logo, :summary, :format, :href, :action, :sort_order)
  end
end
