class ApplicationController < ActionController::Base
  include Pundit
  protect_from_forgery with: :exception
  before_action :set_locale
  before_action :configure_permitted_parameters, if: :devise_controller?

  def set_locale
    I18n.locale = params[:locale] || I18n.default_locale
  end

  def default_url_options
    { locale: I18n.locale }
  end


  protected

  def configure_permitted_parameters
    devise_parameter_sanitizer.permit :sign_up, keys:  [:name, :email, :password, :password_confirmation]
    devise_parameter_sanitizer.permit :account_update, keys:  [:name,  :password, :password_confirmation]
  end


  def require_admin!
    authorize :dashboard, :admin?
  end
end
